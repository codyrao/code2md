package impl

import (
	"code2md/models"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type PythonParser struct {
	*BaseParser
}

func NewPythonParser() *PythonParser {
	return &PythonParser{
		BaseParser: &BaseParser{
			Language:          "python",
			Extensions:        []string{".py", ".pyw"},
			DefaultExtensions: []string{".py", ".pyw"},
		},
	}
}

func (p *PythonParser) Parse(filePath string, content []byte) (*models.FileInfo, error) {
	contentStr := string(content)

	fileInfo := &models.FileInfo{
		Path:         filePath,
		Language:     "python",
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
		if line != "" && !strings.HasPrefix(line, "#") && !strings.HasPrefix(line, "\"\"\"") && !strings.HasPrefix(line, "'''") {
			codeLines++
		}
	}
	fileInfo.CodeLines = codeLines

	fileInfo.Imports = p.extractImports(contentStr)

	classes := p.extractClasses(contentStr, contentStr, filePath)
	fileInfo.Elements = append(fileInfo.Elements, classes...)

	functions := p.extractFunctions(contentStr, contentStr, filePath)
	fileInfo.Elements = append(fileInfo.Elements, functions...)

	variables := p.extractVariables(contentStr, contentStr, filePath)
	fileInfo.Elements = append(fileInfo.Elements, variables...)

	constants := p.extractConstants(contentStr, contentStr, filePath)
	fileInfo.Elements = append(fileInfo.Elements, constants...)

	moduleName := extractModuleName(filePath)
	for _, elem := range fileInfo.Elements {
		if elem.IsExported {
			fileInfo.Exports = append(fileInfo.Exports, moduleName+"."+elem.Name)
		}
	}

	return fileInfo, nil
}

func (p *PythonParser) extractImports(content string) []models.ImportInfo {
	imports := make([]models.ImportInfo, 0)

	matches := regexp.MustCompile(`(?m)^(?:from\s+([a-zA-Z0-9_.]+)\s+)?import\s+([^\n]+)`).FindAllStringSubmatch(content, -1)

	for _, match := range matches {
		if len(match) < 3 {
			continue
		}

		module := match[1]
		importedItems := match[2]

		items := strings.Split(importedItems, ",")
		for _, item := range items {
			item = strings.TrimSpace(item)
			if item == "" {
				continue
			}

			if strings.Contains(item, " as ") {
				parts := strings.Split(item, " as ")
				itemName := strings.TrimSpace(parts[0])
				alias := strings.TrimSpace(parts[1])
				if itemName == "*" {
					imports = append(imports, models.ImportInfo{
						Path:  module,
						Alias: alias,
					})
				} else {
					imports = append(imports, models.ImportInfo{
						Path:  module + "." + itemName,
						Alias: alias,
					})
				}
			} else if item == "*" {
				imports = append(imports, models.ImportInfo{
					Path: module,
				})
			} else {
				fullPath := item
				if module != "" {
					fullPath = module + "." + item
				}
				imports = append(imports, models.ImportInfo{
					Path: fullPath,
				})
			}
		}
	}

	return imports
}

func (p *PythonParser) extractClasses(content string, fullContent string, filePath string) []*models.CodeElement {
	classes := make([]*models.CodeElement, 0)

	classRegex := regexp.MustCompile(`(?:^|\n)\s*(?:class\s+(\w+)(?:\s*\(\s*([A-Z]\w*(?:\s*,\s*[A-Z]\w*)*)\s*\))?\s*:)`)
	classMatches := classRegex.FindAllStringSubmatch(content, -1)

	for _, match := range classMatches {
		if len(match) < 2 {
			continue
		}

		name := match[1]
		bases := ""
		if len(match) > 2 {
			bases = match[2]
		}

		classPattern := regexp.MustCompile(`(?:^|\n)\s*class\s+` + regexp.QuoteMeta(name) + `(?:[^:]*):`)
		classStartMatch := classPattern.FindStringIndex(content)
		if classStartMatch == nil {
			continue
		}

		classContent := content[classStartMatch[0]:]
		endIdx := findMatchingIndent(classContent)

		startLine := strings.Count(content[:classStartMatch[0]], "\n") + 1
		endLine := strings.Count(content[:classStartMatch[0]+endIdx], "\n") + 1

		docComment := extractPythonDocstring(classContent)
		description := extractDescriptionFromPythonDocstring(docComment)

		methods := p.extractClassMethods(classContent)
		classVars := p.extractClassVariables(classContent)

		elem := &models.CodeElement{
			ID:          generateID(),
			Name:        name,
			FullName:    name,
			Type:        models.ElementTypeClass,
			Language:    "python",
			FilePath:    filePath,
			StartLine:   startLine,
			EndLine:     endLine,
			Visibility:  models.VisibilityPublic,
			Description: description,
			DocComment:  docComment,
			LineCount:   endLine - startLine + 1,
			Children:    append(classVars, methods...),
		}

		if bases != "" {
			baseList := strings.Split(bases, ",")
			for i := range baseList {
				baseList[i] = strings.TrimSpace(baseList[i])
			}
			elem.Extends = baseList
		}

		classes = append(classes, elem)

		for _, methodName := range methods {
			methodElem := p.createPythonMethodElement(methodName, classContent, name, filePath, startLine)
			if methodElem != nil {
				classes = append(classes, methodElem)
			}
		}
	}

	return classes
}

func (p *PythonParser) extractClassMethods(classContent string) []string {
	methods := make([]string, 0)

	methodRegex := regexp.MustCompile(`(?m)^\s+def\s+(\w+)\s*\(`)
	matches := methodRegex.FindAllStringSubmatch(classContent, -1)

	for _, match := range matches {
		if len(match) > 1 {
			methodName := match[1]
			if methodName != "__init__" && methodName != "__str__" && methodName != "__repr__" {
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

func (p *PythonParser) extractClassVariables(classContent string) []string {
	variables := make([]string, 0)

	varRegex := regexp.MustCompile(`(?m)^\s+(\w+)\s*=\s*[^=\n]+`)
	matches := varRegex.FindAllStringSubmatch(classContent, -1)

	keywords := map[string]bool{
		"def": true, "class": true, "if": true, "elif": true, "else": true,
		"for": true, "while": true, "try": true, "except": true, "finally": true,
		"with": true, "return": true, "raise": true, "pass": true, "break": true,
		"continue": true, "import": true, "from": true, "as": true, "and": true,
		"or": true, "not": true, "in": true, "is": true, "lambda": true,
		"True": true, "False": true, "None": true, "self": true, "cls": true,
	}

	for _, match := range matches {
		if len(match) > 1 {
			varName := match[1]
			if !keywords[varName] {
				found := false
				for _, v := range variables {
					if v == varName {
						found = true
						break
					}
				}
				if !found {
					variables = append(variables, varName)
				}
			}
		}
	}

	return variables
}

func (p *PythonParser) createPythonMethodElement(methodName string, classContent string, className string, filePath string, classStartLine int) *models.CodeElement {
	methodRegex := regexp.MustCompile(`(?m)^\s+def\s+` + regexp.QuoteMeta(methodName) + `\s*\([^)]*\)\s*:`)
	match := methodRegex.FindString(classContent)
	if match == "" {
		return nil
	}

	idx := strings.Index(classContent, match)
	methodContent := classContent[idx:]
	endIdx := findMatchingIndent(methodContent)

	startLine := classStartLine + strings.Count(classContent[:idx], "\n")
	endLine := classStartLine + strings.Count(classContent[:idx+endIdx], "\n")

	parameters := p.extractPythonParameters(methodContent)

	return &models.CodeElement{
		ID:         generateID(),
		Name:       methodName,
		FullName:   className + "." + methodName,
		Type:       models.ElementTypeMethod,
		Language:   "python",
		FilePath:   filePath,
		StartLine:  startLine,
		EndLine:    endLine,
		Visibility: models.VisibilityPublic,
		Parent:     className,
		Parameters: parameters,
		LineCount:  endLine - startLine + 1,
	}
}

func (p *PythonParser) extractPythonParameters(funcContent string) []models.Parameter {
	params := make([]models.Parameter, 0)

	paramRegex := regexp.MustCompile(`def\s+\w+\s*\(([^)]*)\)`)
	match := paramRegex.FindStringSubmatch(funcContent)
	if len(match) < 2 {
		return params
	}

	paramStr := match[1]
	if strings.TrimSpace(paramStr) == "" {
		return params
	}

	paramParts := parsePythonParams(paramStr)
	for _, part := range paramParts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		defaultValue := strings.Contains(part, "=")
		parts := strings.SplitN(part, "=", 2)
		paramName := strings.TrimSpace(parts[0])
		paramType := "Any"

		if strings.Contains(paramName, ":") {
			typeParts := strings.SplitN(paramName, ":", 2)
			paramName = strings.TrimSpace(typeParts[0])
			paramType = strings.TrimSpace(typeParts[1])
		}

		isVarArgs := strings.HasPrefix(paramName, "*")
		isKwargs := strings.HasPrefix(paramName, "**")

		if isVarArgs {
			paramName = strings.TrimPrefix(paramName, "*")
			paramType = "list"
		} else if isKwargs {
			paramName = strings.TrimPrefix(paramName, "**")
			paramType = "dict"
		}

		params = append(params, models.Parameter{
			Name:       paramName,
			Type:       paramType,
			IsOptional: defaultValue || isVarArgs || isKwargs,
		})
	}

	return params
}

func parsePythonParams(paramsStr string) []string {
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

		if ch == '[' || ch == '(' || ch == '{' {
			depth++
			current.WriteRune(ch)
			continue
		}

		if ch == ']' || ch == ')' || ch == '}' {
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

func (p *PythonParser) extractFunctions(content string, fullContent string, filePath string) []*models.CodeElement {
	functions := make([]*models.CodeElement, 0)

	funcRegex := regexp.MustCompile(`(?m)^(?:async\s+)?def\s+(\w+)\s*\([^)]*\)\s*:`)
	funcMatches := funcRegex.FindAllStringSubmatch(content, -1)

	for _, match := range funcMatches {
		if len(match) < 2 {
			continue
		}

		name := match[1]
		if name == "__init__" || name == "__str__" || name == "__repr__" {
			continue
		}

		funcPattern := regexp.MustCompile(`(?m)^(?:async\s+)?def\s+` + regexp.QuoteMeta(name) + `\s*\([^)]*\)\s*:`)
		funcStartMatch := funcPattern.FindStringIndex(content)
		if funcStartMatch == nil {
			continue
		}

		funcContent := content[funcStartMatch[0]:]
		endIdx := findMatchingIndent(funcContent)

		startLine := strings.Count(content[:funcStartMatch[0]], "\n") + 1
		endLine := strings.Count(content[:funcStartMatch[0]+endIdx], "\n") + 1

		docComment := extractPythonDocstring(funcContent)
		description := extractDescriptionFromPythonDocstring(docComment)

		parameters := p.extractPythonParameters(funcContent)

		returnType := extractPythonReturnType(funcContent)

		elem := &models.CodeElement{
			ID:          generateID(),
			Name:        name,
			FullName:    name,
			Type:        models.ElementTypeFunction,
			Language:    "python",
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

	return functions
}

func (p *PythonParser) extractVariables(content string, fullContent string, filePath string) []*models.CodeElement {
	variables := make([]*models.CodeElement, 0)

	varRegex := regexp.MustCompile(`(?m)^(\w+)\s*=\s*([^\n]+)`)
	varMatches := varRegex.FindAllStringSubmatch(content, -1)

	keywords := map[string]bool{
		"def": true, "class": true, "if": true, "elif": true, "else": true,
		"for": true, "while": true, "try": true, "except": true, "finally": true,
		"with": true, "return": true, "raise": true, "pass": true, "break": true,
		"continue": true, "import": true, "from": true, "as": true, "and": true,
		"or": true, "not": true, "in": true, "is": true, "lambda": true,
		"True": true, "False": true, "None": true,
	}

	for _, match := range varMatches {
		if len(match) > 2 {
			name := match[1]
			value := strings.TrimSpace(match[2])

			if keywords[name] {
				continue
			}

			varPattern := regexp.MustCompile(`(?m)^` + regexp.QuoteMeta(name) + `\s*=`)
			varMatch := varPattern.FindStringIndex(content)
			if varMatch == nil {
				continue
			}

			startLine := strings.Count(content[:varMatch[0]], "\n") + 1

			dataType := inferPythonType(value)

			elem := &models.CodeElement{
				ID:         generateID(),
				Name:       name,
				FullName:   name,
				Type:       models.ElementTypeVariable,
				Language:   "python",
				FilePath:   filePath,
				StartLine:  startLine,
				EndLine:    startLine,
				Visibility: models.VisibilityPublic,
				Value:      value,
				DataType:   dataType,
			}

			variables = append(variables, elem)
		}
	}

	return variables
}

func (p *PythonParser) extractConstants(content string, fullContent string, filePath string) []*models.CodeElement {
	constants := make([]*models.CodeElement, 0)

	constRegex := regexp.MustCompile(`(?m)^([A-Z][A-Z0-9_]*)\s*=\s*([^\n]+)`)
	constMatches := constRegex.FindAllStringSubmatch(content, -1)

	for _, match := range constMatches {
		if len(match) > 2 {
			name := match[1]
			value := strings.TrimSpace(match[2])

			constPattern := regexp.MustCompile(`(?m)^` + regexp.QuoteMeta(name) + `\s*=`)
			constMatch := constPattern.FindStringIndex(content)
			if constMatch == nil {
				continue
			}

			startLine := strings.Count(content[:constMatch[0]], "\n") + 1

			elem := &models.CodeElement{
				ID:         generateID(),
				Name:       name,
				FullName:   name,
				Type:       models.ElementTypeConst,
				Language:   "python",
				FilePath:   filePath,
				StartLine:  startLine,
				EndLine:    startLine,
				Visibility: models.VisibilityPublic,
				Value:      value,
				DataType:   inferPythonType(value),
			}

			constants = append(constants, elem)
		}
	}

	return constants
}

func findMatchingIndent(content string) int {
	lines := strings.Split(content, "\n")
	if len(lines) < 2 {
		return len(content)
	}

	baseIndent := 0
	if len(lines) > 1 {
		secondLine := lines[1]
		baseIndent = len(secondLine) - len(strings.TrimLeft(secondLine, " \t"))
	}

	pos := 0
	for i, line := range lines[1:] {
		if i == 0 {
			continue
		}
		trimmed := strings.TrimLeft(line, " \t")
		if trimmed == "" {
			pos += len(line) + 1
			continue
		}

		currentIndent := len(line) - len(trimmed)
		if currentIndent <= baseIndent {
			return pos
		}
		pos += len(line) + 1
	}

	return len(content)
}

func extractPythonDocstring(content string) string {
	lines := strings.SplitN(content, "\n", 3)
	if len(lines) < 2 {
		return ""
	}

	firstLine := strings.TrimSpace(lines[0])
	if !strings.HasPrefix(firstLine, "\"\"\"") && !strings.HasPrefix(firstLine, "'''") {
		return ""
	}

	quotePrefix := firstLine[:3]
	endMarker := quotePrefix

	content = strings.TrimPrefix(content, firstLine)

	if strings.HasPrefix(content, endMarker) {
		return firstLine
	}

	if strings.HasPrefix(content, "\n") || strings.HasPrefix(content, " ") {
		idx := strings.Index(content, endMarker)
		if idx != -1 {
			return quotePrefix + content[:idx+3]
		}
	}

	return ""
}

func extractDescriptionFromPythonDocstring(doc string) string {
	if doc == "" {
		return ""
	}

	doc = strings.TrimPrefix(doc, "\"\"\"")
	doc = strings.TrimPrefix(doc, "'''")
	doc = strings.TrimSuffix(doc, "\"\"\"")
	doc = strings.TrimSuffix(doc, "'''")
	doc = strings.TrimSpace(doc)

	lines := strings.Split(doc, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" {
			return line
		}
	}

	return ""
}

func extractModuleName(filePath string) string {
	baseName := filePath[strings.LastIndex(filePath, "/")+1:]
	if idx := strings.LastIndex(baseName, "."); idx != -1 {
		baseName = baseName[:idx]
	}
	return baseName
}

func extractPythonReturnType(funcContent string) string {
	returnTypeRegex := regexp.MustCompile(`->\s*([A-Za-z\[\], \n]+)\s*:`)
	match := returnTypeRegex.FindStringSubmatch(funcContent)
	if len(match) > 1 {
		return strings.TrimSpace(match[1])
	}
	return ""
}

func inferPythonType(value string) string {
	value = strings.TrimSpace(value)

	if value == "True" || value == "False" {
		return "bool"
	}
	if _, err := strconv.Atoi(value); err == nil {
		return "int"
	}
	if _, err := strconv.ParseFloat(value, 64); err == nil {
		return "float"
	}
	if strings.HasPrefix(value, "\"") || strings.HasPrefix(value, "'") {
		return "str"
	}
	if strings.HasPrefix(value, "[") {
		return "list"
	}
	if strings.HasPrefix(value, "{") {
		if strings.Contains(value, ":") {
			return "dict"
		}
		return "set"
	}
	if strings.HasPrefix(value, "(") {
		return "tuple"
	}

	return "Any"
}

func (p *PythonParser) ParseFile(filePath string) (*models.FileInfo, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return p.Parse(filePath, content)
}

func init() {
	_ = true
}
