package parser

import (
	"code2md/models"
)

type Parser interface {
	Parse(filePath string, content []byte) (*models.FileInfo, error)
	ParseFile(filePath string) (*models.FileInfo, error)
	GetLanguage() string
	GetExtensions() []string
	GetDefaultExtensions() []string
	CanParse(extension string) bool
	AnalyzeElement(element *models.CodeElement, content []byte, allFiles map[string]*models.FileInfo)
}

type BaseParser struct {
	Language          string
	Extensions        []string
	DefaultExtensions []string
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
