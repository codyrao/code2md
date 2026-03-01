package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

type Config struct {
	Path          string
	Output        string
	Languages     []string
	Exclude       []string
	Verbose       bool
	IncludeHidden bool
	NoStructure   bool
	NoCode        bool
	NoAPI         bool
	ExcludeTests  bool
}

var defaultConfig = Config{
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

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) AddFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&c.Path, "path", "p", defaultConfig.Path, "Project path to analyze (default: ./)")
	cmd.Flags().StringVarP(&c.Output, "output", "o", defaultConfig.Output, "Output markdown file name")
	cmd.Flags().StringArrayVarP(&c.Languages, "language", "l", defaultConfig.Languages, "Programming languages to include (default: all)")
	cmd.Flags().StringArrayVarP(&c.Exclude, "exclude", "e", defaultConfig.Exclude, "Patterns to exclude (e.g., 'vendor/*', '*.test.go')")
	cmd.Flags().BoolVarP(&c.Verbose, "verbose", "v", defaultConfig.Verbose, "Enable verbose output")
	cmd.Flags().BoolVarP(&c.IncludeHidden, "hidden", "H", defaultConfig.IncludeHidden, "Include hidden files and directories")
	cmd.Flags().BoolVar(&c.NoStructure, "no-structure", defaultConfig.NoStructure, "Skip directory tree structure in output")
	cmd.Flags().BoolVar(&c.NoCode, "no-code", defaultConfig.NoCode, "Skip including full code in output")
	cmd.Flags().BoolVar(&c.NoAPI, "no-api", defaultConfig.NoAPI, "Skip API documentation section")
	cmd.Flags().BoolVar(&c.ExcludeTests, "exclude-tests", defaultConfig.ExcludeTests, "Exclude test files and test functions")
}

func (c *Config) Validate() error {
	path, err := os.Stat(c.Path)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("path does not exist: %s", c.Path)
		}
		return fmt.Errorf("error accessing path: %w", err)
	}

	if !path.IsDir() {
		return fmt.Errorf("path is not a directory: %s", c.Path)
	}

	if c.Output == "" {
		return fmt.Errorf("output file name cannot be empty")
	}

	validLanguages := map[string]bool{
		"go":         true,
		"javascript": true,
		"js":         true,
		"typescript": true,
		"ts":         true,
		"java":       true,
		"python":     true,
		"py":         true,
		"cpp":        true,
		"c++":        true,
		"c":          true,
		"php":        true,
	}

	for _, lang := range c.Languages {
		if !validLanguages[lang] {
			return fmt.Errorf("invalid language: %s (valid: go, javascript, typescript, java, python, cpp, c, php)", lang)
		}
	}

	return nil
}

func (c *Config) GetAbsolutePath() (string, error) {
	absPath, err := filepath.Abs(c.Path)
	if err != nil {
		return "", fmt.Errorf("failed to get absolute path: %w", err)
	}
	return absPath, nil
}

func (c *Config) GetOutputPath() (string, error) {
	outputDir := filepath.Dir(c.Output)
	if outputDir == "." {
		outputDir = c.Path
	}

	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		if err := os.MkdirAll(outputDir, 0755); err != nil {
			return "", fmt.Errorf("failed to create output directory: %w", err)
		}
	}

	return filepath.Abs(c.Output)
}

func (c *Config) ShouldIncludeLanguage(language string) bool {
	if len(c.Languages) == 0 {
		return true
	}

	langLower := strings.ToLower(language)
	for _, lang := range c.Languages {
		if strings.ToLower(lang) == langLower {
			return true
		}
	}
	return false
}

func (c *Config) ShouldExcludePath(path string) bool {
	for _, pattern := range c.Exclude {
		if matchPattern(path, pattern) {
			return true
		}
	}
	return false
}

func matchPattern(path, pattern string) bool {
	if pattern == "" {
		return false
	}

	path = strings.ToLower(path)
	pattern = strings.ToLower(pattern)

	if strings.HasPrefix(pattern, "*") {
		suffix := strings.TrimPrefix(pattern, "*")
		return strings.HasSuffix(path, suffix)
	}

	if strings.HasSuffix(pattern, "*") {
		prefix := strings.TrimSuffix(pattern, "*")
		return strings.HasPrefix(path, prefix)
	}

	return path == pattern
}

func init() {
	_, _ = filepath.Abs("")
}
