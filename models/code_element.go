package models

type ElementType string

const (
	ElementTypeFunction  ElementType = "function"
	ElementTypeVariable  ElementType = "variable"
	ElementTypeClass     ElementType = "class"
	ElementTypeInterface ElementType = "interface"
	ElementTypeStruct    ElementType = "struct"
	ElementTypeMethod    ElementType = "method"
	ElementTypeConst     ElementType = "constant"
	ElementTypeImport    ElementType = "import"
	ElementTypePackage   ElementType = "package"
	ElementTypeEnum      ElementType = "enum"
	ElementTypeNamespace ElementType = "namespace"
	ElementTypeModule    ElementType = "module"
)

type Visibility string

const (
	VisibilityPublic    Visibility = "public"
	VisibilityPrivate   Visibility = "private"
	VisibilityProtected Visibility = "protected"
	VisibilityInternal  Visibility = "internal"
)

type Parameter struct {
	Name       string `json:"name"`
	Type       string `json:"type"`
	Desc       string `json:"description,omitempty"`
	IsOptional bool   `json:"is_optional"`
}

type CodeElement struct {
	ID             string            `json:"id"`
	Name           string            `json:"name"`
	FullName       string            `json:"full_name"`
	Type           ElementType       `json:"type"`
	Language       string            `json:"language"`
	FilePath       string            `json:"file_path"`
	StartLine      int               `json:"start_line"`
	EndLine        int               `json:"end_line"`
	StartColumn    int               `json:"start_column,omitempty"`
	EndColumn      int               `json:"end_column,omitempty"`
	Visibility     Visibility        `json:"visibility"`
	Description    string            `json:"description,omitempty"`
	DocComment     string            `json:"doc_comment,omitempty"`
	Parameters     []Parameter       `json:"parameters,omitempty"`
	ReturnType     string            `json:"return_type,omitempty"`
	ReturnDesc     string            `json:"return_description,omitempty"`
	Value          string            `json:"value,omitempty"`
	DataType       string            `json:"data_type"`
	Extends        []string          `json:"extends,omitempty"`
	Implements     []string          `json:"implements,omitempty"`
	Modifiers      []string          `json:"modifiers,omitempty"`
	Tags           map[string]string `json:"tags,omitempty"`
	Calls          []string          `json:"calls,omitempty"`
	CalledBy       []string          `json:"called_by,omitempty"`
	References     []string          `json:"references,omitempty"`
	ReferencedBy   []string          `json:"referenced_by,omitempty"`
	Children       []string          `json:"children,omitempty"`
	Parent         string            `json:"parent,omitempty"`
	IsExported     bool              `json:"is_exported"`
	Signature      string            `json:"signature,omitempty"`
	Complexity     int               `json:"complexity,omitempty"`
	LineCount      int               `json:"line_count"`
	CommentDensity float64           `json:"comment_density,omitempty"`
}

type FileInfo struct {
	Path         string         `json:"path"`
	Language     string         `json:"language"`
	RelativePath string         `json:"relative_path"`
	Package      string         `json:"package,omitempty"`
	ModuleName   string         `json:"module_name,omitempty"`
	Elements     []*CodeElement `json:"elements"`
	Imports      []ImportInfo   `json:"imports,omitempty"`
	Exports      []string       `json:"exports,omitempty"`
	LineCount    int            `json:"line_count"`
	CodeLines    int            `json:"code_lines"`
	CommentLines int            `json:"comment_lines"`
}

type ImportInfo struct {
	Path         string `json:"path"`
	Alias        string `json:"alias,omitempty"`
	ImportType   string `json:"type,omitempty"`
	IsSideEffect bool   `json:"is_side_effect"`
}

type ProjectInfo struct {
	Name          string            `json:"name"`
	RootPath      string            `json:"root_path"`
	Language      string            `json:"primary_language,omitempty"`
	Languages     []string          `json:"languages"`
	Files         []*FileInfo       `json:"files"`
	AllElements   []*CodeElement    `json:"all_elements"`
	DirectoryTree string            `json:"directory_tree"`
	Summary       *ProjectSummary   `json:"summary"`
	Metadata      map[string]string `json:"metadata,omitempty"`
}

type ProjectSummary struct {
	TotalFiles      int            `json:"total_files"`
	TotalLines      int            `json:"total_lines"`
	TotalCodeLines  int            `json:"total_code_lines"`
	TotalComments   int            `json:"total_comments"`
	LanguageStats   map[string]int `json:"language_stats"`
	ElementStats    map[string]int `json:"element_stats"`
	TopLevelDecls   int            `json:"top_level_declarations"`
	Functions       int            `json:"functions"`
	Classes         int            `json:"classes"`
	Interfaces      int            `json:"interfaces"`
	Structs         int            `json:"structs"`
	Variables       int            `json:"variables"`
	Constants       int            `json:"constants"`
	ImportsCount    int            `json:"imports_count"`
	AverageFileSize int            `json:"average_file_size"`
}

type CallGraph struct {
	Nodes map[string]*CallNode `json:"nodes"`
	Edges []CallEdge           `json:"edges"`
}

type CallNode struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Type     string   `json:"type"`
	File     string   `json:"file"`
	Line     int      `json:"line"`
	Outgoing []string `json:"outgoing"`
	Incoming []string `json:"incoming"`
	Visited  bool     `json:"visited"`
	Depth    int      `json:"depth"`
}

type CallEdge struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Type     string `json:"type"`
	Line     int    `json:"line"`
	CallType string `json:"call_type"`
}

type Relationship struct {
	Source      string `json:"source"`
	Target      string `json:"target"`
	RelType     string `json:"relationship_type"`
	Description string `json:"description"`
	File        string `json:"file"`
	Line        int    `json:"line"`
}
