package parser

import (
	"code2md/internal/parser/impl"
	"code2md/models"
	"strings"
)

type ParserFactory struct {
	parsers     map[string]Parser
	byExtension map[string]Parser
}

var globalFactory *ParserFactory

func NewParserFactory() *ParserFactory {
	if globalFactory == nil {
		globalFactory = &ParserFactory{
			parsers:     make(map[string]Parser),
			byExtension: make(map[string]Parser),
		}
		globalFactory.registerDefaultParsers()
	}
	return globalFactory
}

func (f *ParserFactory) registerDefaultParsers() {
	goParser := impl.NewGoParser()
	f.RegisterParser(goParser)

	jsParser := impl.NewJavaScriptParser()
	f.RegisterParser(jsParser)

	tsParser := impl.NewTypeScriptParser()
	f.RegisterParser(tsParser)

	javaParser := impl.NewJavaParser()
	f.RegisterParser(javaParser)

	pythonParser := impl.NewPythonParser()
	f.RegisterParser(pythonParser)

	cppParser := impl.NewCppParser()
	f.RegisterParser(cppParser)

	phpParser := impl.NewPhpParser()
	f.RegisterParser(phpParser)
}

func (f *ParserFactory) RegisterParser(parser Parser) {
	f.parsers[strings.ToLower(parser.GetLanguage())] = parser
	for _, ext := range parser.GetExtensions() {
		f.byExtension[strings.ToLower(ext)] = parser
	}
	for _, ext := range parser.GetDefaultExtensions() {
		if _, exists := f.byExtension[strings.ToLower(ext)]; !exists {
			f.byExtension[strings.ToLower(ext)] = parser
		}
	}
}

func (f *ParserFactory) GetParser(language string) Parser {
	return f.parsers[strings.ToLower(language)]
}

func (f *ParserFactory) GetParserByExtension(extension string) Parser {
	ext := strings.ToLower(extension)
	if !strings.HasPrefix(ext, ".") {
		ext = "." + ext
	}
	if parser, ok := f.byExtension[ext]; ok {
		return parser
	}
	return nil
}

func (f *ParserFactory) GetAllParsers() []Parser {
	parsers := make([]Parser, 0, len(f.parsers))
	for _, parser := range f.parsers {
		parsers = append(parsers, parser)
	}
	return parsers
}

func (f *ParserFactory) GetAllLanguages() []string {
	languages := make([]string, 0, len(f.parsers))
	for lang := range f.parsers {
		languages = append(languages, lang)
	}
	return languages
}

func (f *ParserFactory) GetAllExtensions() []string {
	extensions := make([]string, 0, len(f.byExtension))
	for ext := range f.byExtension {
		extensions = append(extensions, ext)
	}
	return extensions
}

func (f *ParserFactory) ParseFile(filePath string, content []byte, language string) (*models.FileInfo, error) {
	var parser Parser

	if language != "" {
		parser = f.GetParser(language)
	}

	if parser == nil {
		ext := getFileExtension(filePath)
		parser = f.GetParserByExtension(ext)
	}

	if parser == nil {
		return nil, nil
	}

	return parser.Parse(filePath, content)
}

func getFileExtension(filePath string) string {
	idx := strings.LastIndex(filePath, ".")
	if idx == -1 {
		return ""
	}
	return filePath[idx:]
}

func GetParserFactory() *ParserFactory {
	return NewParserFactory()
}
