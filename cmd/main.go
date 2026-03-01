package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"code2md/config"
	"code2md/internal/analyzer"
	"code2md/internal/generator"
	"code2md/internal/parser"
	"code2md/internal/utils"
	"code2md/models"

	"github.com/spf13/cobra"
)

var (
	version = "1.0.0"
	commit  = "dev"
	date    = time.Now().Format(time.RFC3339)
)

var cfg *config.Config

var rootCmd = &cobra.Command{
	Use:     "code2md",
	Version: version,
	Short:   "code2md - Generate documentation from source code",
	Long: `code2md is a command-line tool that analyzes source code and generates 
comprehensive markdown documentation. It supports multiple programming 
languages including Go, JavaScript, TypeScript, Java, Python, C++, and PHP.

Features:
- Automatic code element extraction (functions, classes, variables, etc.)
- Call graph and dependency analysis
- Multi-language support
- Configurable output options
- Clean markdown documentation generation`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg = &config.Config{
			Path:          "./",
			Output:        "code.md",
			Languages:     []string{},
			Exclude:       []string{},
			Verbose:       false,
			IncludeHidden: false,
			NoStructure:   false,
			NoCode:        false,
			NoAPI:         false,
			ExcludeTests:  true,
		}

		if cmd.Flags().Changed("path") {
			cfg.Path, _ = cmd.Flags().GetString("path")
		}
		if cmd.Flags().Changed("output") {
			cfg.Output, _ = cmd.Flags().GetString("output")
		}
		if cmd.Flags().Changed("language") {
			cfg.Languages, _ = cmd.Flags().GetStringArray("language")
		}
		if cmd.Flags().Changed("exclude") {
			cfg.Exclude, _ = cmd.Flags().GetStringArray("exclude")
		}
		if cmd.Flags().Changed("verbose") {
			cfg.Verbose, _ = cmd.Flags().GetBool("verbose")
		}
		if cmd.Flags().Changed("hidden") {
			cfg.IncludeHidden, _ = cmd.Flags().GetBool("hidden")
		}
		if cmd.Flags().Changed("no-structure") {
			cfg.NoStructure, _ = cmd.Flags().GetBool("no-structure")
		}
		if cmd.Flags().Changed("no-code") {
			cfg.NoCode, _ = cmd.Flags().GetBool("no-code")
		}
		if cmd.Flags().Changed("no-api") {
			cfg.NoAPI, _ = cmd.Flags().GetBool("no-api")
		}
		if cmd.Flags().Changed("exclude-tests") {
			cfg.ExcludeTests, _ = cmd.Flags().GetBool("exclude-tests")
		}

		if err := cfg.Validate(); err != nil {
			return err
		}

		return run()
	},
}

func init() {
	cfg := config.NewConfig()
	cfg.AddFlags(rootCmd)
}

func main() {
	rootCmd.PersistentFlags().BoolP("help", "h", false, "Help for code2md")

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	startTime := time.Now()

	absPath, err := cfg.GetAbsolutePath()
	if err != nil {
		return fmt.Errorf("failed to get absolute path: %w", err)
	}

	if cfg.Verbose {
		log.Printf("Analyzing project at: %s", absPath)
	}

	files, err := utils.ScanDirectory(absPath, cfg.Exclude, cfg.ExcludeTests)
	if err != nil {
		return fmt.Errorf("failed to scan directory: %w", err)
	}

	if len(files) == 0 {
		log.Println("No code files found in the specified directory.")
		return nil
	}

	if cfg.Verbose {
		log.Printf("Found %d code files", len(files))
	}

	parserFactory := parser.GetParserFactory()
	projectInfo := createProjectInfo(absPath)

	for _, filePath := range files {
		if cfg.Verbose {
			log.Printf("Processing: %s", filepath.Base(filePath))
		}

		content, err := os.ReadFile(filePath)
		if err != nil {
			if cfg.Verbose {
				log.Printf("Warning: failed to read file %s: %v", filePath, err)
			}
			continue
		}

		ext := utils.GetFileExtension(filePath)
		lang := utils.GetLanguageByExtension(ext)

		if lang == "" || !cfg.ShouldIncludeLanguage(lang) {
			continue
		}

		fileParser := parserFactory.GetParserByExtension(ext)
		if fileParser == nil {
			if cfg.Verbose {
				log.Printf("Warning: no parser found for extension %s", ext)
			}
			continue
		}

		fileInfo, err := fileParser.Parse(filePath, content)
		if err != nil {
			if cfg.Verbose {
				log.Printf("Warning: failed to parse %s: %v", filePath, err)
			}
			continue
		}

		if fileInfo != nil {
			relPath, err := filepath.Rel(absPath, filePath)
			if err != nil {
				relPath = filePath
			}
			fileInfo.RelativePath = relPath

			lines := strings.Split(string(content), "\n")
			fileInfo.LineCount = len(lines)
			fileInfo.CodeLines = utils.CountCodeLines(content)

			projectInfo.Files = append(projectInfo.Files, fileInfo)
			projectInfo.AllElements = append(projectInfo.AllElements, fileInfo.Elements...)
		}
	}

	if len(projectInfo.Files) == 0 {
		log.Println("No files were successfully parsed.")
		return nil
	}

	if cfg.Verbose {
		log.Printf("Parsed %d files", len(projectInfo.Files))
	}

	projectInfo = calculateProjectSummary(projectInfo)

	analyzerInstance := analyzer.NewAnalyzer()
	callGraph := analyzerInstance.Analyze(projectInfo)

	if cfg.Verbose {
		log.Printf("Analyzed %d code elements", len(projectInfo.AllElements))
	}

	tree := utils.GenerateDirectoryTree(absPath, "", cfg.ExcludeTests)
	projectInfo.DirectoryTree = tree

	gen := generator.NewGenerator()
	docContent, err := gen.Generate(projectInfo, callGraph)
	if err != nil {
		return fmt.Errorf("failed to generate documentation: %w", err)
	}

	outputPath, err := cfg.GetOutputPath()
	if err != nil {
		return fmt.Errorf("failed to get output path: %w", err)
	}

	if err := os.WriteFile(outputPath, docContent, 0644); err != nil {
		return fmt.Errorf("failed to write output file: %w", err)
	}

	elapsed := time.Since(startTime)
	fmt.Printf("Documentation generated successfully!\n")
	fmt.Printf("Output file: %s\n", outputPath)
	fmt.Printf("Processing time: %v\n", elapsed)
	fmt.Printf("Files processed: %d\n", len(projectInfo.Files))
	fmt.Printf("Code elements found: %d\n", len(projectInfo.AllElements))

	return nil
}

func createProjectInfo(rootPath string) *models.ProjectInfo {
	projectName := utils.GetProjectName(rootPath)

	return &models.ProjectInfo{
		Name:        projectName,
		RootPath:    rootPath,
		Languages:   make([]string, 0),
		Files:       make([]*models.FileInfo, 0),
		AllElements: make([]*models.CodeElement, 0),
		Metadata: map[string]string{
			"generated_at": time.Now().Format(time.RFC3339),
			"version":      version,
			"commit":       commit,
		},
	}
}

func calculateProjectSummary(project *models.ProjectInfo) *models.ProjectInfo {
	summary := &models.ProjectSummary{
		LanguageStats: make(map[string]int),
		ElementStats:  make(map[string]int),
	}

	summary.TotalFiles = len(project.Files)
	summary.TotalCodeLines = 0
	summary.TotalComments = 0

	languageSet := make(map[string]bool)

	for _, file := range project.Files {
		summary.TotalCodeLines += file.CodeLines
		summary.TotalComments += utils.CountCommentLines([]byte(file.RelativePath))

		summary.LanguageStats[file.Language]++

		lang := strings.ToLower(file.Language)
		if !languageSet[lang] {
			languageSet[lang] = true
			project.Languages = append(project.Languages, file.Language)
		}

		for _, elem := range file.Elements {
			elemType := strings.ToLower(string(elem.Type))
			summary.ElementStats[elemType]++

			switch elem.Type {
			case models.ElementTypeFunction:
				summary.Functions++
			case models.ElementTypeClass, models.ElementTypeStruct:
				summary.Classes++
			case models.ElementTypeInterface:
				summary.Interfaces++
			case models.ElementTypeVariable:
				summary.Variables++
			case models.ElementTypeConst:
				summary.Constants++
			}

			if elem.Type == models.ElementTypeImport {
				summary.ImportsCount++
			}
		}
	}

	if summary.TotalFiles > 0 {
		summary.AverageFileSize = summary.TotalCodeLines / summary.TotalFiles
	}

	project.Summary = summary

	return project
}

func getFileContentLines(filePath string, startLine, endLine int) []string {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil
	}

	lines := strings.Split(string(content), "\n")

	if startLine < 1 {
		startLine = 1
	}
	if endLine > len(lines) {
		endLine = len(lines)
	}

	return lines[startLine-1 : endLine]
}

func formatElementForDisplay(elem *models.CodeElement) string {
	var buf bytes.Buffer

	buf.WriteString(string(elem.Type))
	buf.WriteString(": ")
	buf.WriteString(elem.Name)

	if elem.Description != "" {
		buf.WriteString(" - ")
		buf.WriteString(elem.Description)
	}

	buf.WriteString(" (")
	buf.WriteString(elem.FilePath)
	buf.WriteString(":")
	buf.WriteString(fmt.Sprintf("%d", elem.StartLine))
	buf.WriteString(")")

	return buf.String()
}
