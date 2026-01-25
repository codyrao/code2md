package impl

import (
	"code2md/models"
	"os"
	"regexp"
	"strings"
)

type CppParser struct {
	*BaseParser
}

func NewCppParser() *CppParser {
	return &CppParser{
		BaseParser: &BaseParser{
			Language:          "cpp",
			Extensions:        []string{".cpp", ".cc", ".cxx", ".hpp", ".hh", ".hxx", ".h", ".c"},
			DefaultExtensions: []string{".cpp", ".hpp", ".c", ".h"},
		},
	}
}

func (p *CppParser) Parse(filePath string, content []byte) (*models.FileInfo, error) {
	contentStr := string(content)

	fileInfo := &models.FileInfo{
		Path:         filePath,
		Language:     "cpp",
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

	includes := p.extractIncludes(contentStr)
	fileInfo.Imports = includes

	namespaces := p.extractNamespaces(contentStr, contentStr, filePath)
	fileInfo.Elements = append(fileInfo.Elements, namespaces...)

	classes := p.extractClasses(contentStr, contentStr, filePath)
	fileInfo.Elements = append(fileInfo.Elements, classes...)

	structs := p.extractStructs(contentStr, contentStr, filePath)
	fileInfo.Elements = append(fileInfo.Elements, structs...)

	functions := p.extractFunctions(contentStr, contentStr, filePath)
	fileInfo.Elements = append(fileInfo.Elements, functions...)

	variables := p.extractVariables(contentStr, contentStr, filePath)
	fileInfo.Elements = append(fileInfo.Elements, variables...)

	constants := p.extractConstants(contentStr, contentStr, filePath)
	fileInfo.Elements = append(fileInfo.Elements, constants...)

	enums := p.extractEnums(contentStr, contentStr, filePath)
	fileInfo.Elements = append(fileInfo.Elements, enums...)

	typedefs := p.extractTypedefs(contentStr, contentStr, filePath)
	fileInfo.Elements = append(fileInfo.Elements, typedefs...)

	return fileInfo, nil
}

func (p *CppParser) extractIncludes(content string) []models.ImportInfo {
	imports := make([]models.ImportInfo, 0)

	includeRegex := regexp.MustCompile(`#include\s+(?:<([^>]+)>|"([^"]+)")`)
	matches := includeRegex.FindAllStringSubmatch(content, -1)

	for _, match := range matches {
		includePath := ""
		if len(match) > 1 && match[1] != "" {
			includePath = match[1]
		} else if len(match) > 2 && match[2] != "" {
			includePath = match[2]
		}

		if includePath != "" {
			importType := "system"
			if strings.HasPrefix(includePath, "\"") {
				importType = "local"
			}
			imports = append(imports, models.ImportInfo{
				Path:       includePath,
				ImportType: importType,
			})
		}
	}

	return imports
}

func (p *CppParser) extractNamespaces(content string, fullContent string, filePath string) []*models.CodeElement {
	namespaces := make([]*models.CodeElement, 0)

	namespaceRegex := regexp.MustCompile(`namespace\s+(\w+)\s*\{`)
	namespaceMatches := namespaceRegex.FindAllStringSubmatch(content, -1)

	for _, match := range namespaceMatches {
		if len(match) < 2 {
			continue
		}

		name := match[1]

		idx := strings.Index(content, match[0])
		namespaceContent := content[idx:]
		endIdx := findMatchingBrace(namespaceContent)

		startLine := strings.Count(content[:idx], "\n") + 1
		endLine := strings.Count(content[:idx+endIdx], "\n") + 1

		elem := &models.CodeElement{
			ID:         generateID(),
			Name:       name,
			FullName:   name,
			Type:       models.ElementTypeNamespace,
			Language:   "cpp",
			FilePath:   filePath,
			StartLine:  startLine,
			EndLine:    endLine,
			Visibility: models.VisibilityPublic,
			LineCount:  endLine - startLine + 1,
		}

		namespaces = append(namespaces, elem)
	}

	return namespaces
}

func (p *CppParser) extractClasses(content string, fullContent string, filePath string) []*models.CodeElement {
	classes := make([]*models.CodeElement, 0)

	classRegex := regexp.MustCompile(`(?:class|struct)\s+(\w+)(?:\s*:\s*(?:public|private|protected)\s+(\w+))?(?:\s*,\s*(?:public|private|protected)\s+(\w+))?\s*\{`)
	classMatches := classRegex.FindAllStringSubmatch(content, -1)

	for _, match := range classMatches {
		if len(match) < 2 {
			continue
		}

		name := match[1]
		baseClass := ""
		if len(match) > 2 && match[2] != "" {
			baseClass = match[2]
		}

		isStruct := strings.Contains(match[0], "struct ")
		elemType := models.ElementTypeClass
		if isStruct {
			elemType = models.ElementTypeStruct
		}

		idx := strings.Index(content, match[0])
		classContent := content[idx:]
		endIdx := findMatchingBrace(classContent)

		startLine := strings.Count(content[:idx], "\n") + 1
		endLine := strings.Count(content[:idx+endIdx], "\n") + 1

		docComment := extractCppDocComment(classContent)
		description := extractDescriptionFromCppDoc(docComment)

		methods := p.extractClassMethods(classContent)
		fields := p.extractClassFields(classContent)

		visibility := models.VisibilityPublic
		if !isStruct {
			visibility = models.VisibilityPrivate
		}

		elem := &models.CodeElement{
			ID:          generateID(),
			Name:        name,
			FullName:    name,
			Type:        elemType,
			Language:    "cpp",
			FilePath:    filePath,
			StartLine:   startLine,
			EndLine:     endLine,
			Visibility:  visibility,
			Description: description,
			DocComment:  docComment,
			LineCount:   endLine - startLine + 1,
			Children:    append(fields, methods...),
		}

		if baseClass != "" {
			elem.Extends = []string{baseClass}
		}

		classes = append(classes, elem)

		for _, methodName := range methods {
			methodElem := p.createCppMethodElement(methodName, classContent, name, filePath, startLine)
			if methodElem != nil {
				classes = append(classes, methodElem)
			}
		}

		for _, fieldName := range fields {
			fieldElem := p.createCppFieldElement(fieldName, classContent, name, filePath, startLine)
			if fieldElem != nil {
				classes = append(classes, fieldElem)
			}
		}
	}

	return classes
}

func (p *CppParser) extractStructs(content string, fullContent string, filePath string) []*models.CodeElement {
	structs := make([]*models.CodeElement, 0)

	structRegex := regexp.MustCompile(`struct\s+(\w+)\s*(?::\s*(?:public|private|protected)\s+(\w+))?\s*\{`)
	structMatches := structRegex.FindAllStringSubmatch(content, -1)

	for _, match := range structMatches {
		if len(match) < 2 {
			continue
		}

		name := match[1]
		baseClass := ""
		if len(match) > 2 && match[2] != "" {
			baseClass = match[2]
		}

		idx := strings.Index(content, match[0])
		structContent := content[idx:]
		endIdx := findMatchingBrace(structContent)

		startLine := strings.Count(content[:idx], "\n") + 1
		endLine := strings.Count(content[:idx+endIdx], "\n") + 1

		fields := p.extractStructFields(structContent)

		elem := &models.CodeElement{
			ID:         generateID(),
			Name:       name,
			FullName:   name,
			Type:       models.ElementTypeStruct,
			Language:   "cpp",
			FilePath:   filePath,
			StartLine:  startLine,
			EndLine:    endLine,
			Visibility: models.VisibilityPublic,
			LineCount:  endLine - startLine + 1,
			Children:   fields,
		}

		if baseClass != "" {
			elem.Extends = []string{baseClass}
		}

		structs = append(structs, elem)

		for _, fieldName := range fields {
			fieldElem := p.createCppFieldElement(fieldName, structContent, name, filePath, startLine)
			if fieldElem != nil {
				structs = append(structs, fieldElem)
			}
		}
	}

	return structs
}

func (p *CppParser) extractClassMethods(classContent string) []string {
	methods := make([]string, 0)

	methodRegex := regexp.MustCompile(`(?:\w+::)?(\w+)\s*\([^)]*\)\s*(?:const)?\s*\{`)
	matches := methodRegex.FindAllStringSubmatch(classContent, -1)

	for _, match := range matches {
		if len(match) > 1 {
			methodName := match[1]
			if methodName != "if" && methodName != "while" && methodName != "for" && methodName != "switch" {
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

func (p *CppParser) extractClassFields(classContent string) []string {
	fields := make([]string, 0)

	fieldRegex := regexp.MustCompile(`(?:\w+\s+)?(\w+)\s+(\w+)\s*(?:=\s*[^;]+)?;`)
	matches := fieldRegex.FindAllStringSubmatch(classContent, -1)

	for _, match := range matches {
		if len(match) > 2 {
			fieldType := match[1]
			fieldName := match[2]
			if fieldType != "if" && fieldType != "while" && fieldType != "for" && fieldType != "switch" && fieldType != "return" && fieldType != "class" && fieldType != "struct" && fieldType != "public" && fieldType != "private" && fieldType != "protected" {
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
	}

	return fields
}

func (p *CppParser) extractStructFields(structContent string) []string {
	fields := make([]string, 0)

	fieldRegex := regexp.MustCompile(`(\w+)\s+(\w+)\s*(?:=\s*[^;]+)?;`)
	matches := fieldRegex.FindAllStringSubmatch(structContent, -1)

	for _, match := range matches {
		if len(match) > 2 {
			fieldType := match[1]
			fieldName := match[2]
			if fieldType != "if" && fieldType != "while" && fieldType != "for" && fieldType != "return" && fieldType != "class" && fieldType != "struct" && fieldType != "public" && fieldType != "private" && fieldType != "protected" {
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
	}

	return fields
}

func (p *CppParser) createCppMethodElement(methodName string, classContent string, className string, filePath string, classStartLine int) *models.CodeElement {
	methodRegex := regexp.MustCompile(`((?:\w+::)?` + regexp.QuoteMeta(methodName) + `\s*\([^)]*\)\s*(?:const)?\s*\{[^}]*\})`)
	match := methodRegex.FindString(classContent)
	if match == "" {
		return nil
	}

	idx := strings.Index(classContent, match)
	methodContent := classContent[idx:]
	endIdx := findMatchingBrace(methodContent)

	startLine := classStartLine + strings.Count(classContent[:idx], "\n")
	endLine := classStartLine + strings.Count(classContent[:idx+endIdx], "\n")

	returnType := extractCppReturnType(match)

	return &models.CodeElement{
		ID:         generateID(),
		Name:       methodName,
		FullName:   className + "::" + methodName,
		Type:       models.ElementTypeMethod,
		Language:   "cpp",
		FilePath:   filePath,
		StartLine:  startLine,
		EndLine:    endLine,
		Visibility: models.VisibilityPublic,
		Parent:     className,
		ReturnType: returnType,
		LineCount:  endLine - startLine + 1,
	}
}

func (p *CppParser) createCppFieldElement(fieldName string, classContent string, className string, filePath string, classStartLine int) *models.CodeElement {
	fieldRegex := regexp.MustCompile(`(\w+)\s+` + regexp.QuoteMeta(fieldName) + `\s*(?:=\s*[^;]+)?;`)
	match := fieldRegex.FindString(classContent)
	if match == "" {
		return nil
	}

	idx := strings.Index(classContent, match)
	startLine := classStartLine + strings.Count(classContent[:idx], "\n")

	dataType := extractFieldTypeCpp(match)

	return &models.CodeElement{
		ID:         generateID(),
		Name:       fieldName,
		FullName:   className + "::" + fieldName,
		Type:       models.ElementTypeVariable,
		Language:   "cpp",
		FilePath:   filePath,
		StartLine:  startLine,
		EndLine:    startLine,
		Visibility: models.VisibilityPublic,
		Parent:     className,
		DataType:   dataType,
	}
}

func (p *CppParser) extractFunctions(content string, fullContent string, filePath string) []*models.CodeElement {
	functions := make([]*models.CodeElement, 0)

	funcRegex := regexp.MustCompile(`^(\w+)\s+(\w+)\s*\([^)]*\)\s*(?:const)?\s*(?:\{|;)`)
	funcMatches := funcRegex.FindAllStringSubmatch(content, -1)

	for _, match := range funcMatches {
		if len(match) < 3 {
			continue
		}

		returnType := match[1]
		funcName := match[2]

		if returnType == "class" || returnType == "struct" || returnType == "if" || returnType == "while" || returnType == "for" || returnType == "switch" {
			continue
		}

		funcPattern := regexp.MustCompile(`(?m)^` + regexp.QuoteMeta(returnType) + `\s+` + regexp.QuoteMeta(funcName) + `\s*\([^)]*\)\s*(?:const)?\s*(?:\{|;)`)
		funcStartMatch := funcPattern.FindStringIndex(content)
		if funcStartMatch == nil {
			continue
		}

		funcContent := content[funcStartMatch[0]:]
		if strings.HasSuffix(funcContent, ";") {
			continue
		}

		endIdx := findMatchingBrace(funcContent)

		startLine := strings.Count(content[:funcStartMatch[0]], "\n") + 1
		endLine := strings.Count(content[:funcStartMatch[0]+endIdx], "\n") + 1

		elem := &models.CodeElement{
			ID:         generateID(),
			Name:       funcName,
			FullName:   funcName,
			Type:       models.ElementTypeFunction,
			Language:   "cpp",
			FilePath:   filePath,
			StartLine:  startLine,
			EndLine:    endLine,
			Visibility: models.VisibilityPublic,
			ReturnType: returnType,
			LineCount:  endLine - startLine + 1,
		}

		functions = append(functions, elem)
	}

	return functions
}

func (p *CppParser) extractVariables(content string, fullContent string, filePath string) []*models.CodeElement {
	variables := make([]*models.CodeElement, 0)

	varRegex := regexp.MustCompile(`(\w+)\s+(\w+)\s*(?:=\s*([^;]+))?;`)
	varMatches := varRegex.FindAllStringSubmatch(content, -1)

	for _, match := range varMatches {
		if len(match) < 3 {
			continue
		}

		varType := match[1]
		varName := match[2]
		value := ""
		if len(match) > 3 {
			value = strings.TrimSpace(match[3])
		}

		if varType == "class" || varType == "struct" || varType == "if" || varType == "while" || varType == "for" || varType == "switch" || varType == "return" || varType == "int" || varType == "void" || varType == "bool" || varType == "double" || varType == "float" || varType == "char" || varType == "long" || varType == "short" || varType == "unsigned" || varType == "const" {
			continue
		}

		varPattern := regexp.MustCompile(`(?m)^` + regexp.QuoteMeta(varType) + `\s+` + regexp.QuoteMeta(varName) + `\s*(?:=|;)`)
		varMatch := varPattern.FindStringIndex(content)
		if varMatch == nil {
			continue
		}

		startLine := strings.Count(content[:varMatch[0]], "\n") + 1

		elem := &models.CodeElement{
			ID:         generateID(),
			Name:       varName,
			FullName:   varName,
			Type:       models.ElementTypeVariable,
			Language:   "cpp",
			FilePath:   filePath,
			StartLine:  startLine,
			EndLine:    startLine,
			Visibility: models.VisibilityPublic,
			DataType:   varType,
			Value:      value,
		}

		variables = append(variables, elem)
	}

	return variables
}

func (p *CppParser) extractConstants(content string, fullContent string, filePath string) []*models.CodeElement {
	constants := make([]*models.CodeElement, 0)

	constRegex := regexp.MustCompile(`const\s+(\w+)\s+(\w+)\s*(?:=\s*([^;]+))?;`)
	constMatches := constRegex.FindAllStringSubmatch(content, -1)

	for _, match := range constMatches {
		if len(match) > 3 {
			constType := match[1]
			constName := match[2]
			value := strings.TrimSpace(match[3])

			constPattern := regexp.MustCompile(`const\s+` + regexp.QuoteMeta(constType) + `\s+` + regexp.QuoteMeta(constName))
			constMatch := constPattern.FindStringIndex(content)
			if constMatch == nil {
				continue
			}

			startLine := strings.Count(content[:constMatch[0]], "\n") + 1

			elem := &models.CodeElement{
				ID:         generateID(),
				Name:       constName,
				FullName:   constName,
				Type:       models.ElementTypeConst,
				Language:   "cpp",
				FilePath:   filePath,
				StartLine:  startLine,
				EndLine:    startLine,
				Visibility: models.VisibilityPublic,
				DataType:   constType,
				Value:      value,
			}

			constants = append(constants, elem)
		}
	}

	return constants
}

func (p *CppParser) extractEnums(content string, fullContent string, filePath string) []*models.CodeElement {
	enums := make([]*models.CodeElement, 0)

	enumRegex := regexp.MustCompile(`enum\s+(?:class\s+)?(\w+)\s*\{([^}]*)\}`)
	enumMatches := enumRegex.FindAllStringSubmatch(content, -1)

	for _, match := range enumMatches {
		if len(match) > 2 {
			name := match[1]
			enumBody := match[2]

			idx := strings.Index(content, match[0])
			startLine := strings.Count(content[:idx], "\n") + 1
			endLine := strings.Count(content[:idx+len(match[0])], "\n") + 1

			values := make([]string, 0)
			valueRegex := regexp.MustCompile(`(\w+)(?:\s*=\s*[^,]+)?`)
			valueMatches := valueRegex.FindAllStringSubmatch(enumBody, -1)
			for _, vmatch := range valueMatches {
				if len(vmatch) > 1 {
					values = append(values, vmatch[1])
				}
			}

			elem := &models.CodeElement{
				ID:         generateID(),
				Name:       name,
				FullName:   name,
				Type:       models.ElementTypeEnum,
				Language:   "cpp",
				FilePath:   filePath,
				StartLine:  startLine,
				EndLine:    endLine,
				Visibility: models.VisibilityPublic,
				LineCount:  endLine - startLine + 1,
				Children:   values,
			}

			enums = append(enums, elem)
		}
	}

	return enums
}

func (p *CppParser) extractTypedefs(content string, fullContent string, filePath string) []*models.CodeElement {
	typedefs := make([]*models.CodeElement, 0)

	typedefRegex := regexp.MustCompile(`typedef\s+(\w+)\s+(\w+)\s*;`)
	typedefMatches := typedefRegex.FindAllStringSubmatch(content, -1)

	for _, match := range typedefMatches {
		if len(match) > 2 {
			originalType := match[1]
			newName := match[2]

			idx := strings.Index(content, match[0])
			startLine := strings.Count(content[:idx], "\n") + 1

			elem := &models.CodeElement{
				ID:         generateID(),
				Name:       newName,
				FullName:   newName,
				Type:       models.ElementTypeClass,
				Language:   "cpp",
				FilePath:   filePath,
				StartLine:  startLine,
				EndLine:    startLine,
				Visibility: models.VisibilityPublic,
				DataType:   originalType,
			}

			typedefs = append(typedefs, elem)
		}
	}

	return typedefs
}

func extractCppDocComment(content string) string {
	docRegex := regexp.MustCompile(`///[^\n]*|\/\*\*[\s\S]*?\*\/`)
	match := docRegex.FindString(content)
	return match
}

func extractDescriptionFromCppDoc(doc string) string {
	if doc == "" {
		return ""
	}

	descRegex := regexp.MustCompile(`///\s*(?:\\brief\s*)?(.+)`)
	match := descRegex.FindStringSubmatch(doc)
	if len(match) > 1 {
		return strings.TrimSpace(match[1])
	}

	return ""
}

func extractCppReturnType(funcDecl string) string {
	returnType := strings.TrimSpace(funcDecl)
	if idx := strings.Index(returnType, "("); idx > 0 {
		returnType = strings.TrimSpace(returnType[:idx])
	}
	return returnType
}

func extractFieldTypeCpp(fieldDecl string) string {
	parts := strings.Fields(fieldDecl)
	for _, part := range parts {
		if part == "const" || part == "static" || part == "*" || part == "&" {
			continue
		}
		return part
	}
	return "unknown"
}

func (p *CppParser) ParseFile(filePath string) (*models.FileInfo, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return p.Parse(filePath, content)
}

func (p *BaseParser) ParseFile(filePath string) (*models.FileInfo, error) {
	return nil, nil
}
