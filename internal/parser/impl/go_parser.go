package impl

import (
	"code2md/models"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

type BaseParser struct {
	Language          string
	Extensions        []string
	DefaultExtensions []string
}

type GoParser struct {
	*BaseParser
}

func NewGoParser() *GoParser {
	return &GoParser{
		BaseParser: &BaseParser{
			Language:          "go",
			Extensions:        []string{".go"},
			DefaultExtensions: []string{".go"},
		},
	}
}

func (p *GoParser) Parse(filePath string, content []byte) (*models.FileInfo, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, filePath, content, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	fileInfo := &models.FileInfo{
		Path:         filePath,
		Language:     "go",
		RelativePath: filePath,
		Package:      file.Name.Name,
		Elements:     make([]*models.CodeElement, 0),
		Imports:      make([]models.ImportInfo, 0),
		Exports:      make([]string, 0),
	}

	lines := strings.Split(string(content), "\n")
	fileInfo.LineCount = len(lines)

	codeLines := 0
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" && !strings.HasPrefix(line, "//") && !strings.HasPrefix(line, "/*") && !strings.HasPrefix(line, "*") {
			codeLines++
		}
	}
	fileInfo.CodeLines = codeLines

	for _, imp := range file.Imports {
		importInfo := models.ImportInfo{
			Path: strings.Trim(imp.Path.Value, "\""),
		}
		if imp.Name != nil {
			importInfo.Alias = imp.Name.Name
		}
		fileInfo.Imports = append(fileInfo.Imports, importInfo)
	}

	ast.Inspect(file, func(n ast.Node) bool {
		switch node := n.(type) {
		case *ast.FuncDecl:
			elem := p.parseFuncDecl(node, fset, filePath)
			if elem != nil {
				fileInfo.Elements = append(fileInfo.Elements, elem)
			}
		case *ast.GenDecl:
			for _, spec := range node.Specs {
				switch s := spec.(type) {
				case *ast.TypeSpec:
					elem := p.parseTypeSpec(s, fset, filePath)
					if elem != nil {
						fileInfo.Elements = append(fileInfo.Elements, elem)
					}
				case *ast.ValueSpec:
					for _, name := range s.Names {
						elem := p.parseValueSpec(name, s, fset, filePath, node.Tok)
						if elem != nil {
							fileInfo.Elements = append(fileInfo.Elements, elem)
						}
					}
				case *ast.ImportSpec:
				}
			}
		}
		return true
	})

	for _, elem := range fileInfo.Elements {
		if elem.IsExported {
			fileInfo.Exports = append(fileInfo.Exports, elem.Name)
		}
	}

	return fileInfo, nil
}

func (p *GoParser) parseFuncDecl(node *ast.FuncDecl, fset *token.FileSet, filePath string) *models.CodeElement {
	if node.Recv == nil && node.Name.IsExported() == false && node.Name.Name != "main" {
		return nil
	}

	elem := &models.CodeElement{
		ID:         generateID(),
		Name:       node.Name.Name,
		FullName:   node.Name.Name,
		Type:       models.ElementTypeFunction,
		Language:   "go",
		FilePath:   filePath,
		StartLine:  fset.Position(node.Pos()).Line,
		EndLine:    fset.Position(node.End()).Line,
		Visibility: getVisibility(node.Name.Name),
		IsExported: node.Name.IsExported(),
		LineCount:  fset.Position(node.End()).Line - fset.Position(node.Pos()).Line + 1,
	}

	if node.Recv != nil && len(node.Recv.List) > 0 {
		elem.Type = models.ElementTypeMethod
		recvType := p.getTypeExpr(node.Recv.List[0].Type)
		elem.Parent = recvType
	}

	elem.Signature = p.getFuncSignature(node, fset)

	if node.Doc != nil {
		docComment := node.Doc.Text()
		elem.DocComment = docComment
		elem.Description = extractDescription(docComment)
	}

	if node.Type.Params != nil {
		for _, param := range node.Type.Params.List {
			paramType := p.getTypeExpr(param.Type)
			for _, name := range param.Names {
				elem.Parameters = append(elem.Parameters, models.Parameter{
					Name: name.Name,
					Type: paramType,
				})
			}
		}
	}

	if node.Type.Results != nil {
		if len(node.Type.Results.List) > 0 {
			elem.ReturnType = p.getTypeExpr(node.Type.Results.List[0].Type)
		}
	}

	calls := extractFunctionCalls(node.Body, fset, filePath)
	elem.Calls = calls

	return elem
}

func (p *GoParser) parseTypeSpec(node *ast.TypeSpec, fset *token.FileSet, filePath string) *models.CodeElement {
	elem := &models.CodeElement{
		ID:         generateID(),
		Name:       node.Name.Name,
		FullName:   node.Name.Name,
		Language:   "go",
		FilePath:   filePath,
		StartLine:  fset.Position(node.Pos()).Line,
		EndLine:    fset.Position(node.End()).Line,
		Visibility: getVisibility(node.Name.Name),
		IsExported: node.Name.IsExported(),
		LineCount:  fset.Position(node.End()).Line - fset.Position(node.Pos()).Line + 1,
	}

	switch t := node.Type.(type) {
	case *ast.StructType:
		elem.Type = models.ElementTypeStruct
		for _, field := range t.Fields.List {
			if field.Names != nil {
				for _, name := range field.Names {
					fieldType := p.getTypeExpr(field.Type)
					elem.Children = append(elem.Children, name.Name)
					_ = fieldType
				}
			}
		}
	case *ast.InterfaceType:
		elem.Type = models.ElementTypeInterface
		for _, method := range t.Methods.List {
			if method.Names != nil {
				for _, name := range method.Names {
					elem.Children = append(elem.Children, name.Name)
				}
			}
		}
	default:
		elem.Type = models.ElementTypeClass
	}

	if node.Doc != nil {
		elem.DocComment = node.Doc.Text()
		elem.Description = extractDescription(node.Doc.Text())
	}

	return elem
}

func (p *GoParser) parseValueSpec(name *ast.Ident, spec *ast.ValueSpec, fset *token.FileSet, filePath string, tok token.Token) *models.CodeElement {
	elem := &models.CodeElement{
		ID:         generateID(),
		Name:       name.Name,
		FullName:   name.Name,
		Language:   "go",
		FilePath:   filePath,
		StartLine:  fset.Position(name.Pos()).Line,
		EndLine:    fset.Position(name.End()).Line,
		Visibility: getVisibility(name.Name),
		IsExported: name.IsExported(),
	}

	if tok == token.CONST {
		elem.Type = models.ElementTypeConst
		elem.Value = p.getValue(spec)
	} else {
		elem.Type = models.ElementTypeVariable
		if len(spec.Values) > 0 {
			elem.Value = p.getValue(spec)
		}
	}

	if spec.Type != nil {
		elem.DataType = p.getTypeExpr(spec.Type)
	} else if len(spec.Values) > 0 {
		elem.DataType = p.getTypeExpr(spec.Values[0])
	}

	if spec.Doc != nil {
		elem.DocComment = spec.Doc.Text()
		elem.Description = extractDescription(spec.Doc.Text())
	}

	return elem
}

func (p *GoParser) getFuncSignature(node *ast.FuncDecl, fset *token.FileSet) string {
	var params []string
	if node.Type.Params != nil {
		for _, param := range node.Type.Params.List {
			paramType := p.getTypeExpr(param.Type)
			names := make([]string, len(param.Names))
			for i, name := range param.Names {
				names[i] = name.Name
			}
			if len(names) > 0 {
				params = append(params, strings.Join(names, ", ")+" "+paramType)
			} else {
				params = append(params, paramType)
			}
		}
	}

	result := "func"
	if node.Recv != nil {
		recvType := p.getTypeExpr(node.Recv.List[0].Type)
		result += " (" + recvType + ")"
	}
	result += " " + node.Name.Name + "(" + strings.Join(params, ", ") + ")"

	if node.Type.Results != nil {
		var results []string
		for _, res := range node.Type.Results.List {
			results = append(results, p.getTypeExpr(res.Type))
		}
		result += " (" + strings.Join(results, ", ") + ")"
	}

	return result
}

func (p *GoParser) getTypeExpr(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.SelectorExpr:
		return p.getTypeExpr(t.X) + "." + t.Sel.Name
	case *ast.ArrayType:
		return "[]" + p.getTypeExpr(t.Elt)
	case *ast.MapType:
		return "map[" + p.getTypeExpr(t.Key) + "]" + p.getTypeExpr(t.Value)
	case *ast.ChanType:
		return "chan " + p.getTypeExpr(t.Value)
	case *ast.FuncType:
		return "func(...)"
	case *ast.InterfaceType:
		return "interface{}"
	case *ast.StructType:
		return "struct{}"
	case *ast.StarExpr:
		return "*" + p.getTypeExpr(t.X)
	case *ast.Ellipsis:
		return "..." + p.getTypeExpr(t.Elt)
	default:
		return "interface{}"
	}
}

func (p *GoParser) getValue(spec *ast.ValueSpec) string {
	if len(spec.Values) > 0 {
		return getExprString(spec.Values[0])
	}
	return ""
}

func getExprString(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.BasicLit:
		return t.Value
	case *ast.Ident:
		return t.Name
	case *ast.SelectorExpr:
		return getExprString(t.X) + "." + t.Sel.Name
	case *ast.BinaryExpr:
		return getExprString(t.X) + " " + t.Op.String() + " " + getExprString(t.Y)
	case *ast.CallExpr:
		return getExprString(t.Fun) + "(...)"
	case *ast.ParenExpr:
		return "(" + getExprString(t.X) + ")"
	case *ast.UnaryExpr:
		return t.Op.String() + getExprString(t.X)
	default:
		return "..."
	}
}

func extractFunctionCalls(node ast.Node, fset *token.FileSet, filePath string) []string {
	calls := make([]string, 0)
	ast.Inspect(node, func(n ast.Node) bool {
		if call, ok := n.(*ast.CallExpr); ok {
			if ident, ok := call.Fun.(*ast.Ident); ok {
				calls = append(calls, ident.Name)
			} else if sel, ok := call.Fun.(*ast.SelectorExpr); ok {
				if x, ok := sel.X.(*ast.Ident); ok {
					calls = append(calls, x.Name+"."+sel.Sel.Name)
				}
			}
		}
		return true
	})
	return calls
}

func getVisibility(name string) models.Visibility {
	if len(name) > 0 {
		if name[0] >= 'A' && name[0] <= 'Z' {
			return models.VisibilityPublic
		}
		return models.VisibilityPrivate
	}
	return models.VisibilityPublic
}

func extractDescription(doc string) string {
	lines := strings.Split(doc, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		line = strings.TrimPrefix(line, "*")
		line = strings.TrimSpace(line)
		if line != "" && !strings.HasPrefix(line, "@") && !strings.HasPrefix(line, "//") {
			return line
		}
	}
	return ""
}

func generateID() string {
	return randomString(16)
}

func randomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[0]
	}
	return string(b)
}

func (p *BaseParser) GetLanguage() string {
	return p.Language
}

func (p *BaseParser) GetExtensions() []string {
	return p.Extensions
}

func (p *BaseParser) GetDefaultExtensions() []string {
	return p.DefaultExtensions
}

func (p *BaseParser) CanParse(extension string) bool {
	for _, ext := range p.Extensions {
		if ext == extension {
			return true
		}
	}
	for _, ext := range p.DefaultExtensions {
		if ext == extension {
			return true
		}
	}
	return false
}

func (p *BaseParser) AnalyzeElement(element *models.CodeElement, content []byte, allFiles map[string]*models.FileInfo) {
}

func (p *GoParser) ParseFile(filePath string) (*models.FileInfo, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return p.Parse(filePath, content)
}

func init() {
	_ = filepath.SkipDir
}
