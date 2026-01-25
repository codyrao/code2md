package impl

import (
	"code2md/models"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type JavaParser struct {
	*BaseParser
}

func NewJavaParser() *JavaParser {
	return &JavaParser{
		BaseParser: &BaseParser{
			Language:          "java",
			Extensions:        []string{".java"},
			DefaultExtensions: []string{".java"},
		},
	}
}

func (p *JavaParser) Parse(filePath string, content []byte) (*models.FileInfo, error) {
	contentStr := string(content)

	fileInfo := &models.FileInfo{
		Path:         filePath,
		Language:     "java",
		RelativePath: filePath,
		Elements:     make([]*models.CodeElement, 0),
		Imports:      make([]models.ImportInfo, 0),
		Exports:      make([]string, 0),
	}

	lines := strings.Split(contentStr, "\n")
	fileInfo.LineCount = len(lines)

	codeLines := 0
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" && !strings.HasPrefix(line, "//") && !strings.HasPrefix(line, "/*") && !strings.HasPrefix(line, "*") {
			codeLines++
		}
	}
	fileInfo.CodeLines = codeLines

	packageName := p.extractPackage(contentStr)
	fileInfo.Package = packageName

	fileInfo.Imports = p.extractImports(contentStr)

	classes := p.extractClasses(contentStr, contentStr, filePath)
	fileInfo.Elements = append(fileInfo.Elements, classes...)

	interfaces := p.extractInterfaces(contentStr, contentStr, filePath)
	fileInfo.Elements = append(fileInfo.Elements, interfaces...)

	enums := p.extractEnums(contentStr, contentStr, filePath)
	fileInfo.Elements = append(fileInfo.Elements, enums...)

	for _, elem := range fileInfo.Elements {
		if elem.Type == models.ElementTypeClass && elem.Visibility == models.VisibilityPublic {
			className := elem.Name
			if packageName != "" {
				className = packageName + "." + elem.Name
			}
			fileInfo.Exports = append(fileInfo.Exports, className)
		}
	}

	return fileInfo, nil
}

func (p *JavaParser) extractPackage(content string) string {
	packageRegex := regexp.MustCompile(`package\s+([a-zA-Z0-9_.]+)\s*;`)
	match := packageRegex.FindStringSubmatch(content)
	if len(match) > 1 {
		return match[1]
	}
	return ""
}

func (p *JavaParser) extractImports(content string) []models.ImportInfo {
	imports := make([]models.ImportInfo, 0)

	importRegex := regexp.MustCompile(`import\s+(?:static\s+)?([a-zA-Z0-9_.]+)\s*;`)
	matches := importRegex.FindAllStringSubmatch(content, -1)

	for _, match := range matches {
		if len(match) > 1 {
			importPath := match[1]
			importType := "class"
			if strings.HasPrefix(match[0], "import static") {
				importType = "static"
			} else if strings.HasSuffix(importPath, ".*") {
				importType = "package"
			} else if strings.Contains(importPath, ".") {
				lastPart := importPath[strings.LastIndex(importPath, ".")+1:]
				if firstChar := strings.ToLower(lastPart[:1]); firstChar >= "a" && firstChar <= "z" {
					importType = "class"
				}
			}

			imports = append(imports, models.ImportInfo{
				Path:       importPath,
				ImportType: importType,
			})
		}
	}

	return imports
}

func (p *JavaParser) extractClasses(content string, fullContent string, filePath string) []*models.CodeElement {
	classes := make([]*models.CodeElement, 0)

	classRegex := regexp.MustCompile(`(?:public\s+)?(?:abstract\s+)?(?:final\s+)?class\s+(\w+)(?:\s+extends\s+([A-Z]\w+))?(?:\s+implements\s+([^{]+))?\s*\{`)
	classMatches := classRegex.FindAllStringSubmatch(content, -1)

	for _, match := range classMatches {
		if len(match) < 2 {
			continue
		}

		name := match[1]
		extends := ""
		implements := ""
		if len(match) > 2 {
			extends = match[2]
		}
		if len(match) > 3 {
			implements = match[3]
		}

		idx := strings.Index(content, match[0])
		classContent := content[idx:]
		endIdx := findMatchingBrace(classContent)

		startLine := strings.Count(content[:idx], "\n") + 1
		endLine := strings.Count(content[:idx+endIdx], "\n") + 1

		docComment := extractJavadoc(classContent)
		description := extractDescriptionFromJavadoc(docComment)

		visibility := extractVisibility(match[0])
		isAbstract := strings.Contains(match[0], "abstract")
		isFinal := strings.Contains(match[0], "final")

		modifiers := make([]string, 0)
		if visibility != "" {
			modifiers = append(modifiers, string(visibility))
		}
		if isAbstract {
			modifiers = append(modifiers, "abstract")
		}
		if isFinal {
			modifiers = append(modifiers, "final")
		}

		fields := p.extractClassFields(classContent)
		methods := p.extractClassMethods(classContent)

		elem := &models.CodeElement{
			ID:          generateID(),
			Name:        name,
			FullName:    name,
			Type:        models.ElementTypeClass,
			Language:    "java",
			FilePath:    filePath,
			StartLine:   startLine,
			EndLine:     endLine,
			Visibility:  visibility,
			Description: description,
			DocComment:  docComment,
			LineCount:   endLine - startLine + 1,
			Children:    append(fields, methods...),
			Modifiers:   modifiers,
		}

		if extends != "" {
			elem.Extends = []string{extends}
		}
		if implements != "" {
			implList := strings.Split(implements, ",")
			for i := range implList {
				implList[i] = strings.TrimSpace(implList[i])
			}
			elem.Implements = implList
		}

		classes = append(classes, elem)

		for _, methodName := range methods {
			methodElem := p.createMethodElement(methodName, classContent, name, filePath, startLine)
			classes = append(classes, methodElem)
		}

		for _, fieldName := range fields {
			fieldElem := p.createFieldElement(fieldName, classContent, name, filePath, startLine)
			classes = append(classes, fieldElem)
		}
	}

	return classes
}

func (p *JavaParser) extractClassFields(classContent string) []string {
	fields := make([]string, 0)

	fieldRegex := regexp.MustCompile(`(?:public|private|protected)\s+(?:static|final|transient|volatile)?\s*(?:static\s+)?(?:final\s+)?\s*([A-Z]\w+)\s+(\w+)\s*(?:=\s*[^;]+)?;`)
	matches := fieldRegex.FindAllStringSubmatch(classContent, -1)

	for _, match := range matches {
		if len(match) > 2 {
			fieldName := match[2]
			found := false
			for _, f := range fields {
				if f == fieldName {
					found = true
					break
				}
			}
			if !found {
				fields = append(fields, fieldName)
			}
		}
	}

	return fields
}

func (p *JavaParser) extractClassMethods(classContent string) []string {
	methods := make([]string, 0)

	methodRegex := regexp.MustCompile(`(?:public|private|protected)\s+(?:static\s+)?(?:final\s+)?(?:abstract\s+)?\s*(?:void|int|boolean|double|float|long|short|byte|char|[A-Z]\w*)\s+(\w+)\s*\(`)
	matches := methodRegex.FindAllStringSubmatch(classContent, -1)

	skipMethods := map[string]bool{
		"main": true,
	}

	for _, match := range matches {
		if len(match) > 1 {
			methodName := match[1]
			if !skipMethods[methodName] {
				found := false
				for _, m := range methods {
					if m == methodName {
						found = true
						break
					}
				}
				if !found {
					methods = append(methods, methodName)
				}
			}
		}
	}

	return methods
}

func (p *JavaParser) createMethodElement(methodName string, classContent string, className string, filePath string, classStartLine int) *models.CodeElement {
	methodRegex := regexp.MustCompile(`((?:public|private|protected)\s+(?:static\s+)?(?:final\s+)?(?:abstract\s+)?\s*(?:void|int|boolean|double|float|long|short|byte|char|[A-Z]\w*)\s+` + regexp.QuoteMeta(methodName) + `\([^)]*\)\s*\{[^}]*\})`)
	match := methodRegex.FindString(classContent)
	if match == "" {
		return nil
	}

	idx := strings.Index(classContent, match)
	startLine := classStartLine + strings.Count(classContent[:idx], "\n")

	endLine := classStartLine + strings.Count(classContent[:idx+len(match)], "\n")

	visibility := extractVisibility(match)
	isStatic := strings.Contains(match, "static")
	isFinal := strings.Contains(match, "final")
	isAbstract := strings.Contains(match, "abstract")

	returnType := extractReturnType(match)

	modifiers := make([]string, 0)
	if visibility != "" {
		modifiers = append(modifiers, string(visibility))
	}
	if isStatic {
		modifiers = append(modifiers, "static")
	}
	if isFinal {
		modifiers = append(modifiers, "final")
	}
	if isAbstract {
		modifiers = append(modifiers, "abstract")
	}

	parameters := extractParameters(match)

	return &models.CodeElement{
		ID:         generateID(),
		Name:       methodName,
		FullName:   className + "." + methodName,
		Type:       models.ElementTypeMethod,
		Language:   "java",
		FilePath:   filePath,
		StartLine:  startLine,
		EndLine:    endLine,
		Visibility: visibility,
		Parent:     className,
		ReturnType: returnType,
		Parameters: parameters,
		Modifiers:  modifiers,
		LineCount:  endLine - startLine + 1,
	}
}

func (p *JavaParser) createFieldElement(fieldName string, classContent string, className string, filePath string, classStartLine int) *models.CodeElement {
	fieldRegex := regexp.MustCompile(`(?:public|private|protected)\s+(?:static\s+)?(?:final\s+)?\s*([A-Z]\w*)\s+` + regexp.QuoteMeta(fieldName) + `\s*(?:=\s*[^;]+)?;`)
	match := fieldRegex.FindString(classContent)
	if match == "" {
		return nil
	}

	idx := strings.Index(classContent, match)
	startLine := classStartLine + strings.Count(classContent[:idx], "\n")

	visibility := extractVisibility(match)
	isStatic := strings.Contains(match, "static")
	isFinal := strings.Contains(match, "final")

	dataType := extractFieldType(match)

	modifiers := make([]string, 0)
	if visibility != "" {
		modifiers = append(modifiers, string(visibility))
	}
	if isStatic {
		modifiers = append(modifiers, "static")
	}
	if isFinal {
		modifiers = append(modifiers, "final")
	}

	return &models.CodeElement{
		ID:         generateID(),
		Name:       fieldName,
		FullName:   className + "." + fieldName,
		Type:       models.ElementTypeVariable,
		Language:   "java",
		FilePath:   filePath,
		StartLine:  startLine,
		EndLine:    startLine,
		Visibility: visibility,
		Parent:     className,
		DataType:   dataType,
		Modifiers:  modifiers,
	}
}

func (p *JavaParser) extractInterfaces(content string, fullContent string, filePath string) []*models.CodeElement {
	interfaces := make([]*models.CodeElement, 0)

	interfaceRegex := regexp.MustCompile(`(?:public\s+)?interface\s+(\w+)(?:\s+extends\s+([^{]+))?\s*\{`)
	interfaceMatches := interfaceRegex.FindAllStringSubmatch(content, -1)

	for _, match := range interfaceMatches {
		if len(match) < 2 {
			continue
		}

		name := match[1]
		extends := ""
		if len(match) > 2 {
			extends = match[2]
		}

		idx := strings.Index(content, match[0])
		interfaceContent := content[idx:]
		endIdx := findMatchingBrace(interfaceContent)

		startLine := strings.Count(content[:idx], "\n") + 1
		endLine := strings.Count(content[:idx+endIdx], "\n") + 1

		docComment := extractJavadoc(interfaceContent)
		description := extractDescriptionFromJavadoc(docComment)

		methods := p.extractInterfaceMethods(interfaceContent)

		elem := &models.CodeElement{
			ID:          generateID(),
			Name:        name,
			FullName:    name,
			Type:        models.ElementTypeInterface,
			Language:    "java",
			FilePath:    filePath,
			StartLine:   startLine,
			EndLine:     endLine,
			Visibility:  models.VisibilityPublic,
			Description: description,
			DocComment:  docComment,
			LineCount:   endLine - startLine + 1,
			Children:    methods,
		}

		if extends != "" {
			extendsList := strings.Split(extends, ",")
			for i := range extendsList {
				extendsList[i] = strings.TrimSpace(extendsList[i])
			}
			elem.Extends = extendsList
		}

		interfaces = append(interfaces, elem)
	}

	return interfaces
}

func (p *JavaParser) extractInterfaceMethods(interfaceContent string) []string {
	methods := make([]string, 0)

	methodRegex := regexp.MustCompile(`(?:\s*)?(?:public\s+)?(?:abstract\s+)?\s*(?:void|int|boolean|double|float|long|short|byte|char|[A-Z]\w*)\s+(\w+)\s*\(`)
	matches := methodRegex.FindAllStringSubmatch(interfaceContent, -1)

	for _, match := range matches {
		if len(match) > 1 {
			methods = append(methods, match[1])
		}
	}

	return methods
}

func (p *JavaParser) extractEnums(content string, fullContent string, filePath string) []*models.CodeElement {
	enums := make([]*models.CodeElement, 0)

	enumRegex := regexp.MustCompile(`(?:public\s+)?enum\s+(\w+)(?:\s+implements\s+([^{]+))?\s*\{([^}]*)\}`)
	enumMatches := enumRegex.FindAllStringSubmatch(content, -1)

	for _, match := range enumMatches {
		if len(match) > 1 {
			name := match[1]
			implements := ""
			enumBody := ""
			if len(match) > 2 {
				implements = match[2]
			}
			if len(match) > 3 {
				enumBody = match[3]
			}

			idx := strings.Index(content, match[0])
			startLine := strings.Count(content[:idx], "\n") + 1
			endLine := strings.Count(content[:idx+len(match[0])], "\n") + 1

			docComment := extractJavadoc(content[idx:])
			description := extractDescriptionFromJavadoc(docComment)

			values := make([]string, 0)
			valueRegex := regexp.MustCompile(`(\w+)(?:\([^)]*\))?`)
			valueMatches := valueRegex.FindAllStringSubmatch(enumBody, -1)
			for _, vmatch := range valueMatches {
				if len(vmatch) > 1 {
					values = append(values, vmatch[1])
				}
			}

			elem := &models.CodeElement{
				ID:          generateID(),
				Name:        name,
				FullName:    name,
				Type:        models.ElementTypeEnum,
				Language:    "java",
				FilePath:    filePath,
				StartLine:   startLine,
				EndLine:     endLine,
				Visibility:  models.VisibilityPublic,
				Description: description,
				DocComment:  docComment,
				LineCount:   endLine - startLine + 1,
				Children:    values,
			}

			if implements != "" {
				implList := strings.Split(implements, ",")
				for i := range implList {
					implList[i] = strings.TrimSpace(implList[i])
				}
				elem.Implements = implList
			}

			enums = append(enums, elem)
		}
	}

	return enums
}

func extractVisibility(declaration string) models.Visibility {
	if strings.Contains(declaration, "public ") {
		return models.VisibilityPublic
	} else if strings.Contains(declaration, "private ") {
		return models.VisibilityPrivate
	} else if strings.Contains(declaration, "protected ") {
		return models.VisibilityProtected
	}
	return models.VisibilityPublic
}

func extractReturnType(methodDecl string) string {
	parts := strings.Fields(methodDecl)
	for i, part := range parts {
		if part == "static" || part == "final" || part == "abstract" || part == "public" || part == "private" || part == "protected" {
			continue
		}
		if i > 0 && (part == "void" || part == "int" || part == "boolean" || part == "double" || part == "float" || part == "long" || part == "short" || part == "byte" || part == "char") {
			return part
		}
		if i > 0 && len(part) > 0 {
			firstChar := part[0]
			if (firstChar >= 'A' && firstChar <= 'Z') || (firstChar >= 'a' && firstChar <= 'z') {
				nextIdx := i + 1
				for nextIdx < len(parts) && parts[nextIdx] != "(" {
					if nextIdx > i+1 {
						return parts[i]
					}
					nextIdx++
				}
				if nextIdx < len(parts) && parts[nextIdx] == "(" {
					return parts[i]
				}
			}
		}
	}
	return "void"
}

func extractFieldType(fieldDecl string) string {
	parts := strings.Fields(fieldDecl)
	for _, part := range parts {
		if part == "public" || part == "private" || part == "protected" || part == "static" || part == "final" || part == "transient" || part == "volatile" {
			continue
		}
		return part
	}
	return "Object"
}

func extractParameters(methodDecl string) []models.Parameter {
	params := make([]models.Parameter, 0)

	paramRegex := regexp.MustCompile(`\(([^)]*)\)`)
	match := paramRegex.FindStringSubmatch(methodDecl)
	if len(match) < 2 {
		return params
	}

	paramStr := match[1]
	if strings.TrimSpace(paramStr) == "" {
		return params
	}

	paramParts := strings.Split(paramStr, ",")
	for _, part := range paramParts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		spaceIdx := strings.LastIndex(part, " ")
		if spaceIdx > 0 {
			paramType := strings.TrimSpace(part[:spaceIdx])
			paramName := strings.TrimSpace(part[spaceIdx:])
			params = append(params, models.Parameter{
				Name: paramName,
				Type: paramType,
			})
		}
	}

	return params
}

func extractJavadoc(content string) string {
	javadocRegex := regexp.MustCompile(`/\*\*[\s\S]*?\*/`)
	match := javadocRegex.FindString(content)
	return match
}

func extractDescriptionFromJavadoc(doc string) string {
	if doc == "" {
		return ""
	}

	descRegex := regexp.MustCompile(`\*\s*@?(?:description|desc)\s*[:\-]?\s*(.+)`)
	match := descRegex.FindStringSubmatch(doc)
	if len(match) > 1 {
		return strings.TrimSpace(match[1])
	}

	firstLineRegex := regexp.MustCompile(`/\*\*\s*\n\s*\*\s*(.+)`)
	match = firstLineRegex.FindStringSubmatch(doc)
	if len(match) > 1 {
		return strings.TrimSpace(match[1])
	}

	return ""
}

func (p *JavaParser) ParseFile(filePath string) (*models.FileInfo, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return p.Parse(filePath, content)
}

func init() {
	_ = strconv.IntSize
}
