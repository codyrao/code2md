package impl

import (
	"code2md/models"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type JavaScriptParser struct {
	*BaseParser
}

func NewJavaScriptParser() *JavaScriptParser {
	return &JavaScriptParser{
		BaseParser: &BaseParser{
			Language:          "javascript",
			Extensions:        []string{".js", ".jsx", ".mjs"},
			DefaultExtensions: []string{".js", ".jsx"},
		},
	}
}

func (p *JavaScriptParser) Parse(filePath string, content []byte) (*models.FileInfo, error) {
	contentStr := string(content)

	fileInfo := &models.FileInfo{
		Path:         filePath,
		Language:     "javascript",
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

	functions := p.extractFunctions(contentStr, contentStr, filePath)
	fileInfo.Elements = append(fileInfo.Elements, functions...)

	classes := p.extractClasses(contentStr, contentStr, filePath)
	fileInfo.Elements = append(fileInfo.Elements, classes...)

	variables := p.extractVariables(contentStr, contentStr, filePath)
	fileInfo.Elements = append(fileInfo.Elements, variables...)

	constants := p.extractConstants(contentStr, contentStr, filePath)
	fileInfo.Elements = append(fileInfo.Elements, constants...)

	return fileInfo, nil
}

func (p *JavaScriptParser) extractImports(content string) []models.ImportInfo {
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

	sideEffectRegex := regexp.MustCompile(`import\s+['"]([^'"]+)['"]`)
	sideEffectMatches := sideEffectRegex.FindAllStringSubmatch(content, -1)
	for _, match := range sideEffectMatches {
		if len(match) > 1 {
			imports = append(imports, models.ImportInfo{
				Path:         match[1],
				ImportType:   "side-effect",
				IsSideEffect: true,
			})
		}
	}

	return imports
}

func (p *JavaScriptParser) extractExports(content string) []string {
	exports := make([]string, 0)

	namedExportRegex := regexp.MustCompile(`export\s+(?:const|let|var|function|class|interface|type)\s+(\w+)`)
	matches := namedExportRegex.FindAllStringSubmatch(content, -1)
	for _, match := range matches {
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

func (p *JavaScriptParser) extractFunctions(content string, fullContent string, filePath string) []*models.CodeElement {
	functions := make([]*models.CodeElement, 0)

	functionRegex := regexp.MustCompile(`(?:function\s+(\w+)\s*\([^)]*\)\s*\{|const\s+(\w+)\s*=\s*(?:async\s*)?function\s*\([^)]*\)\s*\{|(\w+)\s*:\s*(?:async\s*)?function\s*\([^)]*\)\s*\{)`)
	funcMatches := functionRegex.FindAllStringSubmatch(content, -1)

	for _, match := range funcMatches {
		name := ""
		for i := 1; i < len(match); i++ {
			if match[i] != "" {
				name = match[i]
				break
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
		description := extractDescriptionFromJSDoc(docComment)

		calls := p.extractFunctionCalls(funcContent)

		elem := &models.CodeElement{
			ID:          generateID(),
			Name:        name,
			FullName:    name,
			Type:        models.ElementTypeFunction,
			Language:    "javascript",
			FilePath:    filePath,
			StartLine:   startLine,
			EndLine:     endLine,
			Visibility:  models.VisibilityPublic,
			Description: description,
			DocComment:  docComment,
			LineCount:   endLine - startLine + 1,
			Calls:       calls,
		}

		functions = append(functions, elem)
	}

	arrowFuncRegex := regexp.MustCompile(`(?:const|let|var)\s+(\w+)\s*=\s*(?:async\s*)?\([^)]*\)\s*=>`)
	arrowMatches := arrowFuncRegex.FindAllStringSubmatch(content, -1)
	for _, match := range arrowMatches {
		if len(match) > 1 {
			name := match[1]
			idx := strings.Index(content, match[0])
			funcContent := content[idx:]
			endIdx := findMatchingBrace(funcContent)

			startLine := strings.Count(content[:idx], "\n") + 1
			endLine := strings.Count(content[:idx+endIdx], "\n") + 1

			docComment := extractJSDoc(funcContent)
			description := extractDescriptionFromJSDoc(docComment)

			calls := p.extractFunctionCalls(funcContent)

			elem := &models.CodeElement{
				ID:          generateID(),
				Name:        name,
				FullName:    name,
				Type:        models.ElementTypeFunction,
				Language:    "javascript",
				FilePath:    filePath,
				StartLine:   startLine,
				EndLine:     endLine,
				Visibility:  models.VisibilityPublic,
				Description: description,
				DocComment:  docComment,
				LineCount:   endLine - startLine + 1,
				Calls:       calls,
			}

			functions = append(functions, elem)
		}
	}

	return functions
}

func (p *JavaScriptParser) extractClasses(content string, fullContent string, filePath string) []*models.CodeElement {
	classes := make([]*models.CodeElement, 0)

	classRegex := regexp.MustCompile(`class\s+(\w+)(?:\s+extends\s+(\w+))?\s*\{`)
	classMatches := classRegex.FindAllStringSubmatch(content, -1)

	for _, match := range classMatches {
		if len(match) < 2 {
			continue
		}

		name := match[1]
		extends := ""
		if len(match) > 2 {
			extends = match[2]
		}

		idx := strings.Index(content, match[0])
		classContent := content[idx:]
		endIdx := findMatchingBrace(classContent)

		startLine := strings.Count(content[:idx], "\n") + 1
		endLine := strings.Count(content[:idx+endIdx], "\n") + 1

		docComment := extractJSDoc(classContent)
		description := extractDescriptionFromJSDoc(docComment)

		methods := p.extractClassMethods(classContent, filePath)

		elem := &models.CodeElement{
			ID:          generateID(),
			Name:        name,
			FullName:    name,
			Type:        models.ElementTypeClass,
			Language:    "javascript",
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
			elem.Extends = []string{extends}
		}

		classes = append(classes, elem)

		for _, method := range methods {
			classIdx := strings.Index(classContent, "method "+method)
			if classIdx != -1 {
				methodContent := classContent[classIdx:]
				methodEnd := findMatchingBrace(methodContent)
				if methodEnd > 0 {
					startLine := strings.Count(content[:idx+classIdx], "\n") + 1
					endLine := strings.Count(content[:idx+classIdx+methodEnd], "\n") + 1

					methodElem := &models.CodeElement{
						ID:         generateID(),
						Name:       method,
						FullName:   name + "." + method,
						Type:       models.ElementTypeMethod,
						Language:   "javascript",
						FilePath:   filePath,
						StartLine:  startLine,
						EndLine:    endLine,
						Visibility: models.VisibilityPublic,
						Parent:     name,
						LineCount:  endLine - startLine + 1,
					}
					classes = append(classes, methodElem)
				}
			}
		}
	}

	return classes
}

func (p *JavaScriptParser) extractClassMethods(classContent string, filePath string) []string {
	methods := make([]string, 0)

	methodRegex := regexp.MustCompile(`(?:async\s+)?(\w+)\s*\([^)]*\)\s*(?:async\s*)?\{`)
	matches := methodRegex.FindAllStringSubmatch(classContent, -1)

	skipWords := map[string]bool{
		"constructor": true, "if": true, "else": true, "for": true, "while": true,
		"switch": true, "try": true, "catch": true, "finally": true, "return": true,
		"throw": true, "new": true, "typeof": true, "instanceof": true,
	}

	for _, match := range matches {
		if len(match) > 1 {
			name := match[1]
			if !skipWords[name] {
				methods = append(methods, name)
			}
		}
	}

	return methods
}

func (p *JavaScriptParser) extractVariables(content string, fullContent string, filePath string) []*models.CodeElement {
	variables := make([]*models.CodeElement, 0)

	varRegex := regexp.MustCompile(`(?:var|let)\s+(\w+)\s*=\s*([^;]+)`)
	varMatches := varRegex.FindAllStringSubmatch(content, -1)

	for _, match := range varMatches {
		if len(match) > 2 {
			name := match[1]
			value := strings.TrimSpace(match[2])

			idx := strings.Index(content, match[0])
			startLine := strings.Count(content[:idx], "\n") + 1

			elem := &models.CodeElement{
				ID:         generateID(),
				Name:       name,
				FullName:   name,
				Type:       models.ElementTypeVariable,
				Language:   "javascript",
				FilePath:   filePath,
				StartLine:  startLine,
				EndLine:    startLine,
				Visibility: getVisibilityJS(name),
				Value:      value,
				DataType:   inferType(value),
			}

			variables = append(variables, elem)
		}
	}

	return variables
}

func (p *JavaScriptParser) extractConstants(content string, fullContent string, filePath string) []*models.CodeElement {
	constants := make([]*models.CodeElement, 0)

	constRegex := regexp.MustCompile(`const\s+(\w+)\s*=\s*([^;]+)`)
	constMatches := constRegex.FindAllStringSubmatch(content, -1)

	for _, match := range constMatches {
		if len(match) > 2 {
			name := match[1]
			value := strings.TrimSpace(match[2])

			idx := strings.Index(content, match[0])
			startLine := strings.Count(content[:idx], "\n") + 1

			elem := &models.CodeElement{
				ID:         generateID(),
				Name:       name,
				FullName:   name,
				Type:       models.ElementTypeConst,
				Language:   "javascript",
				FilePath:   filePath,
				StartLine:  startLine,
				EndLine:    startLine,
				Visibility: getVisibilityJS(name),
				Value:      value,
				DataType:   inferType(value),
			}

			constants = append(constants, elem)
		}
	}

	return constants
}

func (p *JavaScriptParser) extractFunctionCalls(funcContent string) []string {
	calls := make([]string, 0)

	callRegex := regexp.MustCompile(`(\w+)\s*\(`)
	matches := callRegex.FindAllStringSubmatch(funcContent, -1)

	knownFuncs := map[string]bool{
		"if": true, "while": true, "for": true, "switch": true, "return": true,
		"throw": true, "new": true, "console": true, "require": true, "import": true,
	}

	for _, match := range matches {
		if len(match) > 1 {
			name := match[1]
			if !knownFuncs[name] {
				found := false
				for _, c := range calls {
					if c == name {
						found = true
						break
					}
				}
				if !found {
					calls = append(calls, name)
				}
			}
		}
	}

	return calls
}

func findMatchingBrace(content string) int {
	braceCount := 0
	inString := false
	stringChar := '"'
	escaped := false

	for i, ch := range content {
		if escaped {
			escaped = false
			continue
		}

		if ch == '\\' {
			escaped = true
			continue
		}

		if !inString && (ch == '"' || ch == '\'' || ch == '`') {
			inString = true
			stringChar = ch
			continue
		}

		if inString && ch == stringChar {
			inString = false
			continue
		}

		if !inString {
			if ch == '{' {
				braceCount++
			} else if ch == '}' {
				braceCount--
				if braceCount == 0 {
					return i + 1
				}
			}
		}
	}

	return len(content)
}

func extractJSDoc(content string) string {
	jsdocRegex := regexp.MustCompile(`/\*\*[\s\S]*?\*/`)
	match := jsdocRegex.FindString(content)
	return match
}

func extractDescriptionFromJSDoc(doc string) string {
	if doc == "" {
		return ""
	}

	descRegex := regexp.MustCompile(`\s*\*\s*@?(?:description|desc)\s*[:\-]?\s*(.+)`)
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

func getVisibilityJS(name string) models.Visibility {
	if len(name) > 0 && name[0] >= 'A' && name[0] <= 'Z' {
		return models.VisibilityPublic
	}
	return models.VisibilityPrivate
}

func inferType(value string) string {
	value = strings.TrimSpace(value)

	if value == "true" || value == "false" {
		return "boolean"
	}
	if _, err := strconv.Atoi(value); err == nil {
		return "number"
	}
	if _, err := strconv.ParseFloat(value, 64); err == nil {
		return "number"
	}
	if strings.HasPrefix(value, "\"") || strings.HasPrefix(value, "'") || strings.HasPrefix(value, "`") {
		return "string"
	}
	if strings.HasPrefix(value, "[") {
		return "array"
	}
	if strings.HasPrefix(value, "{") {
		return "object"
	}
	if value == "null" || value == "undefined" {
		return "null/undefined"
	}
	if strings.Contains(value, "=>") {
		return "function"
	}

	return "unknown"
}

func (p *JavaScriptParser) ParseFile(filePath string) (*models.FileInfo, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return p.Parse(filePath, content)
}
