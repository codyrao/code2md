package impl

import (
	"code2md/models"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type TypeScriptParser struct {
	*BaseParser
}

func NewTypeScriptParser() *TypeScriptParser {
	return &TypeScriptParser{
		BaseParser: &BaseParser{
			Language:          "typescript",
			Extensions:        []string{".ts", ".tsx"},
			DefaultExtensions: []string{".ts", ".tsx"},
		},
	}
}

func (p *TypeScriptParser) Parse(filePath string, content []byte) (*models.FileInfo, error) {
	contentStr := string(content)

	fileInfo := &models.FileInfo{
		Path:         filePath,
		Language:     "typescript",
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

	fileInfo.Imports = p.extractImports(contentStr)
	fileInfo.Exports = p.extractExports(contentStr)

	interfaces := p.extractInterfaces(contentStr, contentStr, filePath)
	fileInfo.Elements = append(fileInfo.Elements, interfaces...)

	types := p.extractTypeAliases(contentStr, contentStr, filePath)
	fileInfo.Elements = append(fileInfo.Elements, types...)

	functions := p.extractFunctions(contentStr, contentStr, filePath)
	fileInfo.Elements = append(fileInfo.Elements, functions...)

	classes := p.extractClasses(contentStr, contentStr, filePath)
	fileInfo.Elements = append(fileInfo.Elements, classes...)

	variables := p.extractVariables(contentStr, contentStr, filePath)
	fileInfo.Elements = append(fileInfo.Elements, variables...)

	constants := p.extractConstants(contentStr, contentStr, filePath)
	fileInfo.Elements = append(fileInfo.Elements, constants...)

	enums := p.extractEnums(contentStr, contentStr, filePath)
	fileInfo.Elements = append(fileInfo.Elements, enums...)

	return fileInfo, nil
}

func (p *TypeScriptParser) extractImports(content string) []models.ImportInfo {
	imports := make([]models.ImportInfo, 0)

	importRegex := regexp.MustCompile(`import\s+(?:\{[^}]*\}|\*\s+as\s+\w+|\w+)\s+from\s+['"]([^'"]+)['"]`)
	matches := importRegex.FindAllStringSubmatch(content, -1)
	for _, match := range matches {
		if len(match) > 1 {
			imports = append(imports, models.ImportInfo{
				Path: match[1],
			})
		}
	}

	typeImportRegex := regexp.MustCompile(`import\s+type\s+\{[^}]*\}\s+from\s+['"]([^'"]+)['"]`)
	typeMatches := typeImportRegex.FindAllStringSubmatch(content, -1)
	for _, match := range typeMatches {
		if len(match) > 1 {
			imports = append(imports, models.ImportInfo{
				Path:       match[1],
				ImportType: "type",
			})
		}
	}

	return imports
}

func (p *TypeScriptParser) extractExports(content string) []string {
	exports := make([]string, 0)

	namedExportRegex := regexp.MustCompile(`export\s+(?:const|let|var|function|class|interface|type|enum)\s+(\w+)`)
	matches := namedExportRegex.FindAllStringSubmatch(content, -1)
	for _, match := range matches {
		if len(match) > 1 {
			exports = append(exports, match[1])
		}
	}

	interfaceExportRegex := regexp.MustCompile(`export\s+interface\s+(\w+)`)
	interfaceMatches := interfaceExportRegex.FindAllStringSubmatch(content, -1)
	for _, match := range interfaceMatches {
		if len(match) > 1 {
			exports = append(exports, match[1])
		}
	}

	defaultExportRegex := regexp.MustCompile(`export\s+default\s+(\w+)`)
	defaultMatches := defaultExportRegex.FindAllStringSubmatch(content, -1)
	for _, match := range defaultMatches {
		if len(match) > 1 {
			exports = append(exports, match[1])
		}
	}

	return exports
}

func (p *TypeScriptParser) extractInterfaces(content string, fullContent string, filePath string) []*models.CodeElement {
	interfaces := make([]*models.CodeElement, 0)

	interfaceRegex := regexp.MustCompile(`interface\s+(\w+)(?:\s+extends\s+([^{]+))?\s*\{`)
	interfaceMatches := interfaceRegex.FindAllStringSubmatch(content, -1)

	for _, match := range interfaceMatches {
		if len(match) < 2 {
			continue
		}

		name := match[1]
		extends := ""
		if len(match) > 2 {
			extends = strings.TrimSpace(match[2])
		}

		idx := strings.Index(content, match[0])
		interfaceContent := content[idx:]
		endIdx := findMatchingBrace(interfaceContent)

		startLine := strings.Count(content[:idx], "\n") + 1
		endLine := strings.Count(content[:idx+endIdx], "\n") + 1

		docComment := extractJSDoc(interfaceContent)
		description := extractDescriptionFromJSDocTS(docComment)

		properties := p.extractInterfaceProperties(interfaceContent)

		elem := &models.CodeElement{
			ID:          generateID(),
			Name:        name,
			FullName:    name,
			Type:        models.ElementTypeInterface,
			Language:    "typescript",
			FilePath:    filePath,
			StartLine:   startLine,
			EndLine:     endLine,
			Visibility:  models.VisibilityPublic,
			Description: description,
			DocComment:  docComment,
			LineCount:   endLine - startLine + 1,
			Children:    properties,
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

func (p *TypeScriptParser) extractInterfaceProperties(interfaceContent string) []string {
	properties := make([]string, 0)

	propRegex := regexp.MustCompile(`(?:\?|\s)(\w+)\s*:\s*[^;]+;`)
	matches := propRegex.FindAllStringSubmatch(interfaceContent, -1)

	for _, match := range matches {
		if len(match) > 1 {
			properties = append(properties, match[1])
		}
	}

	methodRegex := regexp.MustCompile(`(\w+)\s*\([^)]*\)\s*:\s*[^;]+;`)
	methodMatches := methodRegex.FindAllStringSubmatch(interfaceContent, -1)

	for _, match := range methodMatches {
		if len(match) > 1 {
			found := false
			for _, prop := range properties {
				if prop == match[1] {
					found = true
					break
				}
			}
			if !found {
				properties = append(properties, match[1])
			}
		}
	}

	return properties
}

func (p *TypeScriptParser) extractTypeAliases(content string, fullContent string, filePath string) []*models.CodeElement {
	types := make([]*models.CodeElement, 0)

	typeRegex := regexp.MustCompile(`type\s+(\w+)\s*=\s*([^{;]+)`)
	typeMatches := typeRegex.FindAllStringSubmatch(content, -1)

	for _, match := range typeMatches {
		if len(match) > 2 {
			name := match[1]
			typeValue := strings.TrimSpace(match[2])

			idx := strings.Index(content, match[0])
			startLine := strings.Count(content[:idx], "\n") + 1

			docComment := extractJSDoc(content[idx:])
			description := extractDescriptionFromJSDocTS(docComment)

			elem := &models.CodeElement{
				ID:          generateID(),
				Name:        name,
				FullName:    name,
				Type:        models.ElementTypeClass,
				Language:    "typescript",
				FilePath:    filePath,
				StartLine:   startLine,
				EndLine:     startLine,
				Visibility:  models.VisibilityPublic,
				Description: description,
				DocComment:  docComment,
				DataType:    typeValue,
			}

			if strings.Contains(typeValue, "|") {
				elem.Tags = map[string]string{"union": typeValue}
			} else if strings.Contains(typeValue, "&") {
				elem.Tags = map[string]string{"intersection": typeValue}
			}

			types = append(types, elem)
		}
	}

	return types
}

func (p *TypeScriptParser) extractFunctions(content string, fullContent string, filePath string) []*models.CodeElement {
	functions := make([]*models.CodeElement, 0)

	functionRegex := regexp.MustCompile(`(?:function\s+(\w+)\s*\([^)]*\)\s*:\s*([^=]+)\s*=|const\s+(\w+)\s*=\s*(?:async\s*)?function\s*\([^)]*\)\s*(?::\s*([^=]+))?\s*=)`)
	funcMatches := functionRegex.FindAllStringSubmatch(content, -1)

	for _, match := range funcMatches {
		name := ""
		returnType := ""
		for i := 1; i < len(match); i++ {
			if match[i] != "" {
				if name == "" {
					name = match[i]
				} else if returnType == "" && (i == 2 || i == 4) {
					returnType = strings.TrimSpace(match[i])
				}
			}
		}
		if name == "" {
			continue
		}

		idx := strings.Index(content, match[0])
		funcContent := content[idx:]
		endIdx := findMatchingBrace(funcContent)

		startLine := strings.Count(content[:idx], "\n") + 1
		endLine := strings.Count(content[:idx+endIdx], "\n") + 1

		docComment := extractJSDoc(funcContent)
		description := extractDescriptionFromJSDocTS(docComment)

		parameters := p.extractFunctionParameters(funcContent)

		elem := &models.CodeElement{
			ID:          generateID(),
			Name:        name,
			FullName:    name,
			Type:        models.ElementTypeFunction,
			Language:    "typescript",
			FilePath:    filePath,
			StartLine:   startLine,
			EndLine:     endLine,
			Visibility:  models.VisibilityPublic,
			Description: description,
			DocComment:  docComment,
			LineCount:   endLine - startLine + 1,
			Parameters:  parameters,
			ReturnType:  returnType,
		}

		functions = append(functions, elem)
	}

	arrowFuncRegex := regexp.MustCompile(`const\s+(\w+)\s*=\s*(?:async\s*)?\([^)]*\)\s*(?::\s*([^=]+))?\s*=>`)
	arrowMatches := arrowFuncRegex.FindAllStringSubmatch(content, -1)
	for _, match := range arrowMatches {
		if len(match) > 1 {
			name := match[1]
			returnType := ""
			if len(match) > 2 {
				returnType = strings.TrimSpace(match[2])
			}

			idx := strings.Index(content, match[0])
			funcContent := content[idx:]
			endIdx := findMatchingBrace(funcContent)

			startLine := strings.Count(content[:idx], "\n") + 1
			endLine := strings.Count(content[:idx+endIdx], "\n") + 1

			docComment := extractJSDoc(funcContent)
			description := extractDescriptionFromJSDocTS(docComment)

			parameters := p.extractFunctionParameters(funcContent)

			elem := &models.CodeElement{
				ID:          generateID(),
				Name:        name,
				FullName:    name,
				Type:        models.ElementTypeFunction,
				Language:    "typescript",
				FilePath:    filePath,
				StartLine:   startLine,
				EndLine:     endLine,
				Visibility:  models.VisibilityPublic,
				Description: description,
				DocComment:  docComment,
				LineCount:   endLine - startLine + 1,
				Parameters:  parameters,
				ReturnType:  returnType,
			}

			functions = append(functions, elem)
		}
	}

	return functions
}

func (p *TypeScriptParser) extractFunctionParameters(funcContent string) []models.Parameter {
	parameters := make([]models.Parameter, 0)

	paramRegex := regexp.MustCompile(`(\w+)\s*:\s*([^,)=]+)`)
	matches := paramRegex.FindAllStringSubmatch(funcContent, -1)

	for _, match := range matches {
		if len(match) > 2 {
			paramType := strings.TrimSpace(match[2])
			isOptional := strings.HasSuffix(paramType, "?")
			if isOptional {
				paramType = strings.TrimSuffix(paramType, "?")
			}

			parameters = append(parameters, models.Parameter{
				Name:       match[1],
				Type:       strings.TrimSpace(paramType),
				IsOptional: isOptional,
			})
		}
	}

	return parameters
}

func (p *TypeScriptParser) extractClasses(content string, fullContent string, filePath string) []*models.CodeElement {
	classes := make([]*models.CodeElement, 0)

	classRegex := regexp.MustCompile(`class\s+(\w+)(?:\s+extends\s+(\w+))?(?:\s+implements\s+([^{]+))?\s*\{`)
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

		docComment := extractJSDoc(classContent)
		description := extractDescriptionFromJSDocTS(docComment)

		properties := p.extractClassProperties(classContent)

		elem := &models.CodeElement{
			ID:          generateID(),
			Name:        name,
			FullName:    name,
			Type:        models.ElementTypeClass,
			Language:    "typescript",
			FilePath:    filePath,
			StartLine:   startLine,
			EndLine:     endLine,
			Visibility:  models.VisibilityPublic,
			Description: description,
			DocComment:  docComment,
			LineCount:   endLine - startLine + 1,
			Children:    properties,
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
	}

	return classes
}

func (p *TypeScriptParser) extractClassProperties(classContent string) []string {
	properties := make([]string, 0)

	propRegex := regexp.MustCompile(`(?:private|public|protected|readonly)?\s*(?:\w+\??)\s*:\s*[^=]+`)
	matches := propRegex.FindAllStringSubmatch(classContent, -1)

	for _, match := range matches {
		propMatch := regexp.MustCompile(`(\w+)\s*:`).FindStringSubmatch(match[0])
		if len(propMatch) > 1 {
			found := false
			for _, prop := range properties {
				if prop == propMatch[1] {
					found = true
					break
				}
			}
			if !found {
				properties = append(properties, propMatch[1])
			}
		}
	}

	return properties
}

func (p *TypeScriptParser) extractVariables(content string, fullContent string, filePath string) []*models.CodeElement {
	variables := make([]*models.CodeElement, 0)

	varRegex := regexp.MustCompile(`(?:var|let)\s+(\w+)\s*:\s*([^=]+)(?:\s*=\s*([^;]+))?`)
	varMatches := varRegex.FindAllStringSubmatch(content, -1)

	for _, match := range varMatches {
		if len(match) > 2 {
			name := match[1]
			varType := strings.TrimSpace(match[2])
			value := ""
			if len(match) > 3 {
				value = strings.TrimSpace(match[3])
			}

			idx := strings.Index(content, match[0])
			startLine := strings.Count(content[:idx], "\n") + 1

			elem := &models.CodeElement{
				ID:         generateID(),
				Name:       name,
				FullName:   name,
				Type:       models.ElementTypeVariable,
				Language:   "typescript",
				FilePath:   filePath,
				StartLine:  startLine,
				EndLine:    startLine,
				Visibility: getVisibilityJS(name),
				DataType:   varType,
				Value:      value,
			}

			variables = append(variables, elem)
		}
	}

	return variables
}

func (p *TypeScriptParser) extractConstants(content string, fullContent string, filePath string) []*models.CodeElement {
	constants := make([]*models.CodeElement, 0)

	constRegex := regexp.MustCompile(`const\s+(\w+)\s*:\s*([^=]+)(?:\s*=\s*([^;]+))?`)
	constMatches := constRegex.FindAllStringSubmatch(content, -1)

	for _, match := range constMatches {
		if len(match) > 2 {
			name := match[1]
			varType := strings.TrimSpace(match[2])
			value := ""
			if len(match) > 3 {
				value = strings.TrimSpace(match[3])
			}

			idx := strings.Index(content, match[0])
			startLine := strings.Count(content[:idx], "\n") + 1

			elem := &models.CodeElement{
				ID:         generateID(),
				Name:       name,
				FullName:   name,
				Type:       models.ElementTypeConst,
				Language:   "typescript",
				FilePath:   filePath,
				StartLine:  startLine,
				EndLine:    startLine,
				Visibility: getVisibilityJS(name),
				DataType:   varType,
				Value:      value,
			}

			constants = append(constants, elem)
		}
	}

	return constants
}

func (p *TypeScriptParser) extractEnums(content string, fullContent string, filePath string) []*models.CodeElement {
	enums := make([]*models.CodeElement, 0)

	enumRegex := regexp.MustCompile(`enum\s+(\w+)\s*\{([^}]*)\}`)
	enumMatches := enumRegex.FindAllStringSubmatch(content, -1)

	for _, match := range enumMatches {
		if len(match) > 2 {
			name := match[1]
			enumBody := match[2]

			idx := strings.Index(content, match[0])
			startLine := strings.Count(content[:idx], "\n") + 1
			endLine := strings.Count(content[:idx+len(match[0])], "\n") + 1

			docComment := extractJSDoc(content[idx:])
			description := extractDescriptionFromJSDocTS(docComment)

			values := make([]string, 0)
			valueRegex := regexp.MustCompile(`(\w+)(?:\s*=\s*[^,]+)?`)
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
				Language:    "typescript",
				FilePath:    filePath,
				StartLine:   startLine,
				EndLine:     endLine,
				Visibility:  models.VisibilityPublic,
				Description: description,
				DocComment:  docComment,
				LineCount:   endLine - startLine + 1,
				Children:    values,
			}

			enums = append(enums, elem)
		}
	}

	return enums
}

func extractDescriptionFromJSDocTS(doc string) string {
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

func (p *TypeScriptParser) ParseFile(filePath string) (*models.FileInfo, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return p.Parse(filePath, content)
}

func init() {
	_ = strconv.IntSize
}
