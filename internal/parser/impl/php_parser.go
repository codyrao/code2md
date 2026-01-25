package impl

import (
	"code2md/models"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type PhpParser struct {
	*BaseParser
}

func NewPhpParser() *PhpParser {
	return &PhpParser{
		BaseParser: &BaseParser{
			Language:          "php",
			Extensions:        []string{".php", ".phtml"},
			DefaultExtensions: []string{".php"},
		},
	}
}

func (p *PhpParser) Parse(filePath string, content []byte) (*models.FileInfo, error) {
	contentStr := string(content)

	fileInfo := &models.FileInfo{
		Path:         filePath,
		Language:     "php",
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
		if line != "" && !strings.HasPrefix(line, "//") && !strings.HasPrefix(line, "#") && !strings.HasPrefix(line, "/*") && !strings.HasPrefix(line, "*") {
			codeLines++
		}
	}
	fileInfo.CodeLines = codeLines

	namespace := p.extractNamespace(contentStr)
	fileInfo.Package = namespace

	fileInfo.Imports = p.extractImports(contentStr)

	classes := p.extractClasses(contentStr, contentStr, filePath)
	fileInfo.Elements = append(fileInfo.Elements, classes...)

	interfaces := p.extractInterfaces(contentStr, contentStr, filePath)
	fileInfo.Elements = append(fileInfo.Elements, interfaces...)

	traits := p.extractTraits(contentStr, contentStr, filePath)
	fileInfo.Elements = append(fileInfo.Elements, traits...)

	functions := p.extractFunctions(contentStr, contentStr, filePath)
	fileInfo.Elements = append(fileInfo.Elements, functions...)

	variables := p.extractVariables(contentStr, contentStr, filePath)
	fileInfo.Elements = append(fileInfo.Elements, variables...)

	constants := p.extractConstants(contentStr, contentStr, filePath)
	fileInfo.Elements = append(fileInfo.Elements, constants...)

	for _, elem := range fileInfo.Elements {
		if elem.Type == models.ElementTypeClass && elem.Visibility == models.VisibilityPublic {
			className := elem.Name
			if namespace != "" {
				className = namespace + "\\" + elem.Name
			}
			fileInfo.Exports = append(fileInfo.Exports, className)
		}
	}

	return fileInfo, nil
}

func (p *PhpParser) extractNamespace(content string) string {
	namespaceRegex := regexp.MustCompile(`namespace\s+([a-zA-Z0-9_\\]+)\s*;`)
	match := namespaceRegex.FindStringSubmatch(content)
	if len(match) > 1 {
		return match[1]
	}
	return ""
}

func (p *PhpParser) extractImports(content string) []models.ImportInfo {
	imports := make([]models.ImportInfo, 0)

	useRegex := regexp.MustCompile(`use\s+([a-zA-Z0-9_\\]+)(?:\s+as\s+(\w+))?\s*;`)
	matches := useRegex.FindAllStringSubmatch(content, -1)

	for _, match := range matches {
		if len(match) > 1 {
			path := match[1]
			alias := ""
			if len(match) > 2 {
				alias = match[2]
			}

			imports = append(imports, models.ImportInfo{
				Path:  path,
				Alias: alias,
			})
		}
	}

	groupUseRegex := regexp.MustCompile(`use\s+([a-zA-Z0-9_\\]+)\s*\\{([^}]+)\}\s*;`)
	groupMatches := groupUseRegex.FindAllStringSubmatch(content, -1)

	for _, match := range groupMatches {
		if len(match) > 2 {
			basePath := match[1]
			importedItems := match[2]
			items := strings.Split(importedItems, ",")
			for _, item := range items {
				item = strings.TrimSpace(item)
				if item == "" {
					continue
				}

				var alias, className string
				if strings.Contains(item, " as ") {
					parts := strings.Split(item, " as ")
					className = strings.TrimSpace(parts[0])
					alias = strings.TrimSpace(parts[1])
				} else {
					className = item
				}

				fullPath := basePath + "\\" + className
				imports = append(imports, models.ImportInfo{
					Path:  fullPath,
					Alias: alias,
				})
			}
		}
	}

	return imports
}

func (p *PhpParser) extractClasses(content string, fullContent string, filePath string) []*models.CodeElement {
	classes := make([]*models.CodeElement, 0)

	classRegex := regexp.MustCompile(`(?:final\s+)?(?:abstract\s+)?class\s+(\w+)(?:\s+extends\s+([A-Z]\w+))?(?:\s+implements\s+([^{]+))?\s*\{`)
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

		docComment := extractPhpDocComment(classContent)
		description := extractDescriptionFromPhpDoc(docComment)

		isAbstract := strings.Contains(match[0], "abstract")
		isFinal := strings.Contains(match[0], "final")

		methods := p.extractClassMethods(classContent)
		properties := p.extractClassProperties(classContent)

		elem := &models.CodeElement{
			ID:          generateID(),
			Name:        name,
			FullName:    name,
			Type:        models.ElementTypeClass,
			Language:    "php",
			FilePath:    filePath,
			StartLine:   startLine,
			EndLine:     endLine,
			Visibility:  models.VisibilityPublic,
			Description: description,
			DocComment:  docComment,
			LineCount:   endLine - startLine + 1,
			Children:    append(properties, methods...),
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

		if isAbstract {
			elem.Modifiers = append(elem.Modifiers, "abstract")
		}
		if isFinal {
			elem.Modifiers = append(elem.Modifiers, "final")
		}

		classes = append(classes, elem)

		for _, methodName := range methods {
			methodElem := p.createPhpMethodElement(methodName, classContent, name, filePath, startLine)
			if methodElem != nil {
				classes = append(classes, methodElem)
			}
		}

		for _, propName := range properties {
			propElem := p.createPhpPropertyElement(propName, classContent, name, filePath, startLine)
			if propElem != nil {
				classes = append(classes, propElem)
			}
		}
	}

	return classes
}

func (p *PhpParser) extractInterfaces(content string, fullContent string, filePath string) []*models.CodeElement {
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
			extends = match[2]
		}

		idx := strings.Index(content, match[0])
		interfaceContent := content[idx:]
		endIdx := findMatchingBrace(interfaceContent)

		startLine := strings.Count(content[:idx], "\n") + 1
		endLine := strings.Count(content[:idx+endIdx], "\n") + 1

		docComment := extractPhpDocComment(interfaceContent)
		description := extractDescriptionFromPhpDoc(docComment)

		methods := p.extractInterfaceMethods(interfaceContent)

		elem := &models.CodeElement{
			ID:          generateID(),
			Name:        name,
			FullName:    name,
			Type:        models.ElementTypeInterface,
			Language:    "php",
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

func (p *PhpParser) extractTraits(content string, fullContent string, filePath string) []*models.CodeElement {
	traits := make([]*models.CodeElement, 0)

	traitRegex := regexp.MustCompile(`trait\s+(\w+)\s*\{`)
	traitMatches := traitRegex.FindAllStringSubmatch(content, -1)

	for _, match := range traitMatches {
		if len(match) < 2 {
			continue
		}

		name := match[1]

		idx := strings.Index(content, match[0])
		traitContent := content[idx:]
		endIdx := findMatchingBrace(traitContent)

		startLine := strings.Count(content[:idx], "\n") + 1
		endLine := strings.Count(content[:idx+endIdx], "\n") + 1

		docComment := extractPhpDocComment(traitContent)
		description := extractDescriptionFromPhpDoc(docComment)

		methods := p.extractClassMethods(traitContent)

		elem := &models.CodeElement{
			ID:          generateID(),
			Name:        name,
			FullName:    name,
			Type:        models.ElementTypeClass,
			Language:    "php",
			FilePath:    filePath,
			StartLine:   startLine,
			EndLine:     endLine,
			Visibility:  models.VisibilityPublic,
			Description: description,
			DocComment:  docComment,
			LineCount:   endLine - startLine + 1,
			Children:    methods,
		}

		traits = append(traits, elem)

		for _, methodName := range methods {
			methodElem := p.createPhpMethodElement(methodName, traitContent, name, filePath, startLine)
			if methodElem != nil {
				traits = append(traits, methodElem)
			}
		}
	}

	return traits
}

func (p *PhpParser) extractClassMethods(classContent string) []string {
	methods := make([]string, 0)

	methodRegex := regexp.MustCompile(`(?:public|private|protected)\s+(?:static\s+)?(?:abstract\s+)?function\s+(\w+)\s*\(`)
	matches := methodRegex.FindAllStringSubmatch(classContent, -1)

	for _, match := range matches {
		if len(match) > 1 {
			methodName := match[1]
			if methodName != "__construct" && methodName != "__destruct" && methodName != "__get" && methodName != "__set" && methodName != "__call" {
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

func (p *PhpParser) extractClassProperties(classContent string) []string {
	properties := make([]string, 0)

	propRegex := regexp.MustCompile(`(?:public|private|protected)\s+(?:static\s+)?(?:\??\s*)?(\$\w+)`)
	matches := propRegex.FindAllStringSubmatch(classContent, -1)

	for _, match := range matches {
		if len(match) > 1 {
			propName := strings.TrimPrefix(match[1], "$")
			found := false
			for _, p := range properties {
				if p == propName {
					found = true
					break
				}
			}
			if !found {
				properties = append(properties, propName)
			}
		}
	}

	return properties
}

func (p *PhpParser) extractInterfaceMethods(interfaceContent string) []string {
	methods := make([]string, 0)

	methodRegex := regexp.MustCompile(`function\s+(\w+)\s*\(`)
	matches := methodRegex.FindAllStringSubmatch(interfaceContent, -1)

	for _, match := range matches {
		if len(match) > 1 {
			methods = append(methods, match[1])
		}
	}

	return methods
}

func (p *PhpParser) createPhpMethodElement(methodName string, classContent string, className string, filePath string, classStartLine int) *models.CodeElement {
	methodRegex := regexp.MustCompile(`((?:public|private|protected)\s+(?:static\s+)?(?:abstract\s+)?function\s+` + regexp.QuoteMeta(methodName) + `\s*\([^)]*\)\s*(?:;|\{[^}]*\}))`)
	match := methodRegex.FindString(classContent)
	if match == "" {
		return nil
	}

	idx := strings.Index(classContent, match)
	methodContent := classContent[idx:]
	endIdx := findMatchingBrace(methodContent)

	startLine := classStartLine + strings.Count(classContent[:idx], "\n")
	endLine := classStartLine + strings.Count(classContent[:idx+endIdx], "\n")

	visibility := extractPhpVisibility(match)
	isStatic := strings.Contains(match, "static")
	isAbstract := strings.Contains(match, "abstract")

	modifiers := make([]string, 0)
	if visibility != "" {
		modifiers = append(modifiers, string(visibility))
	}
	if isStatic {
		modifiers = append(modifiers, "static")
	}
	if isAbstract {
		modifiers = append(modifiers, "abstract")
	}

	parameters := p.extractPhpParameters(methodContent)

	return &models.CodeElement{
		ID:         generateID(),
		Name:       methodName,
		FullName:   className + "::" + methodName,
		Type:       models.ElementTypeMethod,
		Language:   "php",
		FilePath:   filePath,
		StartLine:  startLine,
		EndLine:    endLine,
		Visibility: visibility,
		Parent:     className,
		Parameters: parameters,
		Modifiers:  modifiers,
		LineCount:  endLine - startLine + 1,
	}
}

func (p *PhpParser) createPhpPropertyElement(propName string, classContent string, className string, filePath string, classStartLine int) *models.CodeElement {
	propRegex := regexp.MustCompile(`((?:public|private|protected)\s+(?:static\s+)?` + regexp.QuoteMeta(propName) + `\s*(?:=\s*[^;]+)?;)`)
	match := propRegex.FindString(classContent)
	if match == "" {
		return nil
	}

	idx := strings.Index(classContent, match)
	startLine := classStartLine + strings.Count(classContent[:idx], "\n")

	visibility := extractPhpVisibility(match)
	isStatic := strings.Contains(match, "static")

	modifiers := make([]string, 0)
	if visibility != "" {
		modifiers = append(modifiers, string(visibility))
	}
	if isStatic {
		modifiers = append(modifiers, "static")
	}

	return &models.CodeElement{
		ID:         generateID(),
		Name:       propName,
		FullName:   className + "::$" + propName,
		Type:       models.ElementTypeVariable,
		Language:   "php",
		FilePath:   filePath,
		StartLine:  startLine,
		EndLine:    startLine,
		Visibility: visibility,
		Parent:     className,
		Modifiers:  modifiers,
	}
}

func (p *PhpParser) extractFunctions(content string, fullContent string, filePath string) []*models.CodeElement {
	functions := make([]*models.CodeElement, 0)

	funcRegex := regexp.MustCompile(`function\s+(\w+)\s*\(`)
	funcMatches := funcRegex.FindAllStringSubmatch(content, -1)

	for _, match := range funcMatches {
		if len(match) < 2 {
			continue
		}

		funcName := match[1]

		idx := strings.Index(content, match[0])
		funcContent := content[idx:]
		endIdx := findMatchingBrace(funcContent)

		startLine := strings.Count(content[:idx], "\n") + 1
		endLine := strings.Count(content[:idx+endIdx], "\n") + 1

		docComment := extractPhpDocComment(funcContent)
		description := extractDescriptionFromPhpDoc(docComment)

		parameters := p.extractPhpParameters(funcContent)

		elem := &models.CodeElement{
			ID:          generateID(),
			Name:        funcName,
			FullName:    funcName,
			Type:        models.ElementTypeFunction,
			Language:    "php",
			FilePath:    filePath,
			StartLine:   startLine,
			EndLine:     endLine,
			Visibility:  models.VisibilityPublic,
			Description: description,
			DocComment:  docComment,
			Parameters:  parameters,
			LineCount:   endLine - startLine + 1,
		}

		functions = append(functions, elem)
	}

	return functions
}

func (p *PhpParser) extractVariables(content string, fullContent string, filePath string) []*models.CodeElement {
	variables := make([]*models.CodeElement, 0)

	varRegex := regexp.MustCompile(`(\$\w+)\s*=\s*([^;]+);`)
	varMatches := varRegex.FindAllStringSubmatch(content, -1)

	for _, match := range varMatches {
		if len(match) > 2 {
			varName := strings.TrimPrefix(match[1], "$")
			value := strings.TrimSpace(match[2])

			varPattern := regexp.MustCompile(`(?m)^` + regexp.QuoteMeta(match[1]) + `\s*=`)
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
				Language:   "php",
				FilePath:   filePath,
				StartLine:  startLine,
				EndLine:    startLine,
				Visibility: models.VisibilityPublic,
				Value:      value,
				DataType:   inferPhpType(value),
			}

			variables = append(variables, elem)
		}
	}

	return variables
}

func (p *PhpParser) extractConstants(content string, fullContent string, filePath string) []*models.CodeElement {
	constants := make([]*models.CodeElement, 0)

	constRegex := regexp.MustCompile(`const\s+(\w+)\s*=\s*([^;]+);`)
	constMatches := constRegex.FindAllStringSubmatch(content, -1)

	for _, match := range constMatches {
		if len(match) > 2 {
			constName := match[1]
			value := strings.TrimSpace(match[2])

			constPattern := regexp.MustCompile(`(?m)^const\s+` + regexp.QuoteMeta(constName))
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
				Language:   "php",
				FilePath:   filePath,
				StartLine:  startLine,
				EndLine:    startLine,
				Visibility: models.VisibilityPublic,
				Value:      value,
				DataType:   inferPhpType(value),
			}

			constants = append(constants, elem)
		}
	}

	return constants
}

func (p *PhpParser) extractPhpParameters(funcContent string) []models.Parameter {
	params := make([]models.Parameter, 0)

	paramRegex := regexp.MustCompile(`function\s+\w+\s*\(([^)]*)\)`)
	match := paramRegex.FindStringSubmatch(funcContent)
	if len(match) < 2 {
		return params
	}

	paramStr := match[1]
	if strings.TrimSpace(paramStr) == "" {
		return params
	}

	paramParts := parsePhpParams(paramStr)
	for _, part := range paramParts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		defaultValue := strings.Contains(part, "=")
		parts := strings.SplitN(part, "=", 2)
		paramName := strings.TrimPrefix(strings.TrimSpace(parts[0]), "$")
		paramType := "mixed"

		if strings.Contains(paramName, ":") {
			typeParts := strings.SplitN(paramName, ":", 2)
			paramType = strings.TrimSpace(typeParts[0])
			paramName = strings.TrimSpace(typeParts[1])
		} else if len(parts) > 1 {
			paramType = inferPhpType(parts[1])
		}

		params = append(params, models.Parameter{
			Name:       paramName,
			Type:       paramType,
			IsOptional: defaultValue,
		})
	}

	return params
}

func parsePhpParams(paramsStr string) []string {
	var params []string
	var current strings.Builder
	depth := 0
	inString := false
	var stringChar rune

	for i, ch := range paramsStr {
		if inString {
			if ch == stringChar && (i == 0 || paramsStr[i-1] != '\\') {
				inString = false
			}
			current.WriteRune(ch)
			continue
		}

		if ch == '"' || ch == '\'' {
			inString = true
			stringChar = ch
			current.WriteRune(ch)
			continue
		}

		if ch == '(' || ch == '[' || ch == '{' {
			depth++
			current.WriteRune(ch)
			continue
		}

		if ch == ')' || ch == ']' || ch == '}' {
			depth--
			current.WriteRune(ch)
			continue
		}

		if ch == ',' && depth == 0 {
			params = append(params, current.String())
			current.Reset()
			continue
		}

		current.WriteRune(ch)
	}

	if current.Len() > 0 {
		params = append(params, current.String())
	}

	return params
}

func extractPhpDocComment(content string) string {
	docRegex := regexp.MustCompile(`/\*\*[\s\S]*?\*/`)
	match := docRegex.FindString(content)
	return match
}

func extractDescriptionFromPhpDoc(doc string) string {
	if doc == "" {
		return ""
	}

	descRegex := regexp.MustCompile(`\*\s*@?(?:description|desc|short)\s*[:\-]?\s*(.+)`)
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

func extractPhpVisibility(declaration string) models.Visibility {
	if strings.Contains(declaration, "public ") {
		return models.VisibilityPublic
	} else if strings.Contains(declaration, "private ") {
		return models.VisibilityPrivate
	} else if strings.Contains(declaration, "protected ") {
		return models.VisibilityProtected
	}
	return models.VisibilityPublic
}

func inferPhpType(value string) string {
	value = strings.TrimSpace(value)

	if value == "true" || value == "false" {
		return "bool"
	}
	if _, err := strconv.Atoi(value); err == nil {
		return "int"
	}
	if _, err := strconv.ParseFloat(value, 64); err == nil {
		return "float"
	}
	if strings.HasPrefix(value, "\"") || strings.HasPrefix(value, "'") {
		return "string"
	}
	if strings.HasPrefix(value, "[") || strings.HasPrefix(value, "array(") {
		return "array"
	}
	if strings.HasPrefix(value, "null") || value == "NULL" {
		return "null"
	}

	return "mixed"
}

func (p *PhpParser) ParseFile(filePath string) (*models.FileInfo, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return p.Parse(filePath, content)
}

func init() {
	_ = strconv.IntSize
}
