package generator

import (
	"bytes"
	"code2md/models"
	"fmt"
	"path/filepath"
	"sort"
	"strings"
	"text/template"
)

type Generator struct {
	templates *template.Template
}

type TemplateData struct {
	Project   *models.ProjectInfo
	CallGraph *models.CallGraph
}

func NewGenerator() *Generator {
	g := &Generator{
		templates: template.New("code2md"),
	}
	g.initTemplates()
	return g
}

func (g *Generator) initTemplates() {
	templateText := `
# {{ .Project.Name }} - Code Documentation

## 项目概览

**生成时间**: {{ .Project.Metadata.generated_at }}
**项目路径**: {{ .Project.RootPath }}
**主要语言**: {{ .Project.Language }}

### 语言统计

{{ range $lang, $count := .Project.Summary.LanguageStats }}
- **{{ $lang }}**: {{ $count }} 个文件
{{ end }}

### 项目摘要

| 指标 | 数量 |
|------|------|
| 总文件数 | {{ .Project.Summary.TotalFiles }} |
| 总代码行数 | {{ .Project.Summary.TotalCodeLines }} |
| 总注释行数 | {{ .Project.Summary.TotalComments }} |
| 函数数量 | {{ .Project.Summary.Functions }} |
| 类数量 | {{ .Project.Summary.Classes }} |
| 接口数量 | {{ .Project.Summary.Interfaces }} |
| 变量数量 | {{ .Project.Summary.Variables }} |
| 常量数量 | {{ .Project.Summary.Constants }} |

## 项目结构

{{ .Project.DirectoryTree }}

## 文件清单

{{ range .Project.Files }}
### {{ .RelativePath }}

- **语言**: {{ .Language }}
- **包**: {{ .Package }}
- **代码行数**: {{ .CodeLines }}
- **元素数量**: {{ len .Elements }}

{{ if .Imports }}
#### 导入

{{ range .Imports }}
- ` + "`" + `{{ .Path }}` + "`" + `{{ if .Alias }} (as {{ .Alias }}){{ end }}
{{ end }}
{{ end }}

{{ if .Elements }}
#### 代码元素

{{ range .Elements }}
**{{ .Type | title }}**: ` + "`" + `{{ .Name }}` + "`" + `

- **完整名称**: {{ .FullName }}
- **文件位置**: 行 {{ .StartLine }}-{{ .EndLine }}
- **可见性**: {{ .Visibility }}
{{ if .Description }}- **描述**: {{ .Description }}{{ end }}
{{ if .DataType }}- **类型**: ` + "`" + `{{ .DataType }}` + "`" + `{{ end }}
{{ if .ReturnType }}- **返回值**: ` + "`" + `{{ .ReturnType }}` + "`" + `{{ end }}
{{ if .Parameters }}
- **参数**:
{{ range .Parameters }}  - ` + "`" + `{{ .Name }}` + "`" + `: {{ .Type }}{{ if .IsOptional }} (可选){{ end }}
{{ end }}{{ end }}
{{ if .Extends }}- **继承**: {{ range .Extends }}{{ . }} {{ end }}{{ end }}
{{ if .Implements }}- **实现**: {{ range .Implements }}{{ . }} {{ end }}{{ end }}
{{ if .Value }}- **值**: ` + "`" + `{{ .Value }}` + "`" + `{{ end }}
{{ if .Children }}
- **子元素**: {{ range .Children }}{{ . }} {{ end }}{{ end }}
{{ end }}

{{ end }}
{{ end }}

## API 参考

### 公共接口

{{ range .Project.AllElements }}
{{ if eq .Visibility "public" }}
#### {{ .Type | title }}: {{ .Name }}

{{ if .Signature }}
` + "```" + `{{ .Language }}
{{ .Signature }}
` + "```" + `

{{ end }}
{{ if .Description }}{{ .Description }}

{{ end }}
{{ if .DocComment }}
` + "```" + `
{{ .DocComment }}
` + "```" + `

{{ end }}
{{ if .Parameters }}
**参数**:
{{ range .Parameters }}
- ` + "`" + `{{ .Name }}` + "`" + `: {{ .Type }}{{ if .IsOptional }} (可选){{ end }}
{{ end }}
{{ end }}
{{ if .ReturnType }}
**返回**: ` + "`" + `{{ .ReturnType }}` + "`" + `
{{ end }}
{{ if .Calls }}
**调用**: {{ range .Calls }}{{ . }}() {{ end }}
{{ end }}

---

{{ end }}
{{ end }}

## 调用关系图

{{ range $id, $node := .CallGraph.Nodes }}
### {{ $node.Name }} ({{ $node.Type }})

- **文件**: {{ $node.File }}:{{ $node.Line }}
- **调用深度**: {{ $node.Depth }}
{{ if $node.Incoming }}
**被以下调用**:
{{ range $node.Incoming }}{{ . }} {{ end }}
{{ end }}
{{ if $node.Outgoing }}
**调用以下**:
{{ range $node.Outgoing }}{{ . }} {{ end }}
{{ end }}

{{ end }}
`

	g.templates = template.Must(g.templates.Funcs(
		template.FuncMap{
			"title": strings.Title,
			"add":   func(a, b int) int { return a + b },
		},
	).Parse(templateText))
}

func (g *Generator) Generate(project *models.ProjectInfo, callGraph *models.CallGraph) ([]byte, error) {
	data := TemplateData{
		Project:   project,
		CallGraph: callGraph,
	}

	var buf bytes.Buffer
	err := g.templates.Execute(&buf, data)
	if err != nil {
		return nil, fmt.Errorf("failed to execute template: %w", err)
	}

	return buf.Bytes(), nil
}

func (g *Generator) GenerateSummary(project *models.ProjectInfo) string {
	var sb strings.Builder

	sb.WriteString("# ")
	sb.WriteString(project.Name)
	sb.WriteString(" - 代码文档\n\n")

	sb.WriteString("## 项目概览\n\n")
	sb.WriteString("- **生成时间**: ")
	sb.WriteString(project.Metadata["generated_at"])
	sb.WriteString("\n")
	sb.WriteString("- **项目路径**: ")
	sb.WriteString(project.RootPath)
	sb.WriteString("\n")
	sb.WriteString("- **主要语言**: ")
	sb.WriteString(project.Language)
	sb.WriteString("\n\n")

	sb.WriteString("### 语言统计\n\n")
	langStats := make([]string, 0, len(project.Summary.LanguageStats))
	for lang, count := range project.Summary.LanguageStats {
		langStats = append(langStats, fmt.Sprintf("- **%s**: %d 个文件", lang, count))
	}
	sort.Strings(langStats)
	sb.WriteString(strings.Join(langStats, "\n"))
	sb.WriteString("\n\n")

	sb.WriteString("### 项目摘要\n\n")
	sb.WriteString("| 指标 | 数量 |\n")
	sb.WriteString("|------|------|\n")
	sb.WriteString(fmt.Sprintf("| 总文件数 | %d |\n", project.Summary.TotalFiles))
	sb.WriteString(fmt.Sprintf("| 总代码行数 | %d |\n", project.Summary.TotalCodeLines))
	sb.WriteString(fmt.Sprintf("| 总注释行数 | %d |\n", project.Summary.TotalComments))
	sb.WriteString(fmt.Sprintf("| 函数数量 | %d |\n", project.Summary.Functions))
	sb.WriteString(fmt.Sprintf("| 类数量 | %d |\n", project.Summary.Classes))
	sb.WriteString(fmt.Sprintf("| 接口数量 | %d |\n", project.Summary.Interfaces))
	sb.WriteString(fmt.Sprintf("| 变量数量 | %d |\n", project.Summary.Variables))
	sb.WriteString(fmt.Sprintf("| 常量数量 | %d |\n", project.Summary.Constants))
	sb.WriteString("\n")

	return sb.String()
}

func (g *Generator) GenerateFileSection(file *models.FileInfo) string {
	var sb strings.Builder

	sb.WriteString("### ")
	sb.WriteString(file.RelativePath)
	sb.WriteString("\n\n")

	sb.WriteString("- **语言**: ")
	sb.WriteString(file.Language)
	sb.WriteString("\n")
	sb.WriteString("- **包**: ")
	sb.WriteString(file.Package)
	sb.WriteString("\n")
	sb.WriteString("- **代码行数**: ")
	sb.WriteString(fmt.Sprintf("%d", file.CodeLines))
	sb.WriteString("\n")
	sb.WriteString("- **元素数量**: ")
	sb.WriteString(fmt.Sprintf("%d", len(file.Elements)))
	sb.WriteString("\n\n")

	if len(file.Imports) > 0 {
		sb.WriteString("#### 导入\n\n")
		for _, imp := range file.Imports {
			sb.WriteString("- `")
			sb.WriteString(imp.Path)
			sb.WriteString("`")
			if imp.Alias != "" {
				sb.WriteString(" (as ")
				sb.WriteString(imp.Alias)
				sb.WriteString(")")
			}
			sb.WriteString("\n")
		}
		sb.WriteString("\n")
	}

	return sb.String()
}

func (g *Generator) GenerateElementSection(element *models.CodeElement) string {
	var sb strings.Builder

	sb.WriteString("**")
	sb.WriteString(strings.Title(string(element.Type)))
	sb.WriteString("**: `")
	sb.WriteString(element.Name)
	sb.WriteString("`\n\n")

	sb.WriteString("- **完整名称**: ")
	sb.WriteString(element.FullName)
	sb.WriteString("\n")
	sb.WriteString("- **文件位置**: 行 ")
	sb.WriteString(fmt.Sprintf("%d", element.StartLine))
	sb.WriteString("-")
	sb.WriteString(fmt.Sprintf("%d", element.EndLine))
	sb.WriteString("\n")
	sb.WriteString("- **可见性**: ")
	sb.WriteString(string(element.Visibility))
	sb.WriteString("\n")

	if element.Description != "" {
		sb.WriteString("- **描述**: ")
		sb.WriteString(element.Description)
		sb.WriteString("\n")
	}

	if element.DataType != "" {
		sb.WriteString("- **类型**: `")
		sb.WriteString(element.DataType)
		sb.WriteString("`\n")
	}

	if element.ReturnType != "" {
		sb.WriteString("- **返回值**: `")
		sb.WriteString(element.ReturnType)
		sb.WriteString("`\n")
	}

	if len(element.Parameters) > 0 {
		sb.WriteString("- **参数**:\n")
		for _, param := range element.Parameters {
			sb.WriteString("  - `")
			sb.WriteString(param.Name)
			sb.WriteString("`: ")
			sb.WriteString(param.Type)
			if param.IsOptional {
				sb.WriteString(" (可选)")
			}
			sb.WriteString("\n")
		}
	}

	if len(element.Extends) > 0 {
		sb.WriteString("- **继承**: ")
		sb.WriteString(strings.Join(element.Extends, ", "))
		sb.WriteString("\n")
	}

	if len(element.Implements) > 0 {
		sb.WriteString("- **实现**: ")
		sb.WriteString(strings.Join(element.Implements, ", "))
		sb.WriteString("\n")
	}

	if element.Value != "" {
		sb.WriteString("- **值**: `")
		sb.WriteString(element.Value)
		sb.WriteString("`\n")
	}

	if len(element.Children) > 0 {
		sb.WriteString("- **子元素**: ")
		sb.WriteString(strings.Join(element.Children, ", "))
		sb.WriteString("\n")
	}

	if len(element.Calls) > 0 {
		sb.WriteString("- **调用**: ")
		for i, call := range element.Calls {
			if i > 0 {
				sb.WriteString(", ")
			}
			sb.WriteString("`")
			sb.WriteString(call)
			sb.WriteString("()`")
		}
		sb.WriteString("\n")
	}

	return sb.String()
}

func (g *Generator) GenerateAPISection(elements []*models.CodeElement) string {
	var sb strings.Builder

	for _, elem := range elements {
		if elem.Visibility != models.VisibilityPublic {
			continue
		}

		sb.WriteString("#### ")
		sb.WriteString(strings.Title(string(elem.Type)))
		sb.WriteString(": ")
		sb.WriteString(elem.Name)
		sb.WriteString("\n\n")

		if elem.Signature != "" {
			sb.WriteString("```")
			sb.WriteString(elem.Language)
			sb.WriteString("\n")
			sb.WriteString(elem.Signature)
			sb.WriteString("\n```\n\n")
		}

		if elem.Description != "" {
			sb.WriteString(elem.Description)
			sb.WriteString("\n\n")
		}

		if elem.DocComment != "" {
			sb.WriteString("```\n")
			sb.WriteString(elem.DocComment)
			sb.WriteString("\n```\n\n")
		}

		if len(elem.Parameters) > 0 {
			sb.WriteString("**参数**:\n")
			for _, param := range elem.Parameters {
				sb.WriteString("- `")
				sb.WriteString(param.Name)
				sb.WriteString("`: ")
				sb.WriteString(param.Type)
				if param.IsOptional {
					sb.WriteString(" (可选)")
				}
				sb.WriteString("\n")
			}
			sb.WriteString("\n")
		}

		if elem.ReturnType != "" {
			sb.WriteString("**返回**: `")
			sb.WriteString(elem.ReturnType)
			sb.WriteString("`\n\n")
		}

		if len(elem.Calls) > 0 {
			sb.WriteString("**调用**: ")
			for i, call := range elem.Calls {
				if i > 0 {
					sb.WriteString(", ")
				}
				sb.WriteString("`")
				sb.WriteString(call)
				sb.WriteString("()`")
			}
			sb.WriteString("\n\n")
		}

		sb.WriteString("---\n\n")
	}

	return sb.String()
}

func (g *Generator) GenerateCallGraphSection(callGraph *models.CallGraph) string {
	var sb strings.Builder

	for id, node := range callGraph.Nodes {
		sb.WriteString("### ")
		sb.WriteString(node.Name)
		sb.WriteString(" (")
		sb.WriteString(node.Type)
		sb.WriteString(")\n\n")

		sb.WriteString("- **文件**: ")
		sb.WriteString(node.File)
		sb.WriteString(":")
		sb.WriteString(fmt.Sprintf("%d", node.Line))
		sb.WriteString("\n")
		sb.WriteString("- **调用深度**: ")
		sb.WriteString(fmt.Sprintf("%d", node.Depth))
		sb.WriteString("\n")

		if len(node.Incoming) > 0 {
			sb.WriteString("**被以下调用**:\n")
			for _, incoming := range node.Incoming {
				sb.WriteString("- ")
				sb.WriteString(incoming)
				sb.WriteString("\n")
			}
		}

		if len(node.Outgoing) > 0 {
			sb.WriteString("**调用以下**:\n")
			for _, outgoing := range node.Outgoing {
				sb.WriteString("- ")
				sb.WriteString(outgoing)
				sb.WriteString("\n")
			}
		}

		sb.WriteString("\n")
		_ = id
	}

	return sb.String()
}

func GetRelativePath(absPath, rootPath string) string {
	rel, err := filepath.Rel(rootPath, absPath)
	if err != nil {
		return absPath
	}
	return rel
}

func FormatElementType(elemType models.ElementType) string {
	return strings.Title(string(elemType))
}
