package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type FileInfo struct {
	Path      string
	Extension string
	Content   []byte
	IsDir     bool
}

func IsHiddenFile(path string) bool {
	base := filepath.Base(path)
	return strings.HasPrefix(base, ".") || base == ".git" || base == ".svn" || base == ".hg"
}

func IsExcludedDir(dirName string, excludePatterns []string) bool {
	excludedDirs := []string{
		"node_modules", "vendor", "bower_components", ".git", ".svn", ".hg",
		"__pycache__", ".idea", ".vscode", "dist", "build", "target",
		"out", "obj", "bin", ".gradle", "vendor", "public", "static",
	}

	for _, excluded := range excludedDirs {
		if dirName == excluded {
			return true
		}
	}

	for _, pattern := range excludePatterns {
		if matchPattern(dirName, pattern) {
			return true
		}
	}

	return false
}

func IsCodeFile(extension string, validExtensions map[string]bool) bool {
	return validExtensions[strings.ToLower(extension)]
}

func GetValidExtensions() map[string]bool {
	return map[string]bool{
		".go":    true,
		".js":    true,
		".jsx":   true,
		".ts":    true,
		".tsx":   true,
		".java":  true,
		".py":    true,
		".pyw":   true,
		".cpp":   true,
		".cc":    true,
		".cxx":   true,
		".hpp":   true,
		".hh":    true,
		".hxx":   true,
		".h":     true,
		".c":     true,
		".php":   true,
		".phtml": true,
	}
}

func matchPattern(name, pattern string) bool {
	if pattern == "" {
		return false
	}

	if strings.HasPrefix(pattern, "*") {
		pattern = pattern[1:]
		return strings.HasSuffix(name, pattern)
	}

	if strings.HasSuffix(pattern, "*") {
		pattern = strings.TrimSuffix(pattern, "*")
		return strings.HasPrefix(name, pattern)
	}

	return name == pattern
}

func ScanDirectory(rootPath string, excludePatterns []string) ([]string, error) {
	var files []string

	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(rootPath, path)
		if err != nil {
			return err
		}

		if info.IsDir() {
			dirName := filepath.Base(path)
			if IsExcludedDir(dirName, excludePatterns) {
				return filepath.SkipDir
			}
			return nil
		}

		if IsHiddenFile(relPath) {
			return nil
		}

		ext := filepath.Ext(path)
		if IsCodeFile(ext, GetValidExtensions()) {
			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}

func ReadFileContent(filePath string) ([]byte, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func GetFileSize(filePath string) (int64, error) {
	info, err := os.Stat(filePath)
	if err != nil {
		return 0, err
	}
	return info.Size(), nil
}

func GetFileModificationTime(filePath string) (int64, error) {
	info, err := os.Stat(filePath)
	if err != nil {
		return 0, err
	}
	return info.ModTime().Unix(), nil
}

func CountLines(content []byte) int {
	if len(content) == 0 {
		return 0
	}
	count := 1
	for i := 0; i < len(content); i++ {
		if content[i] == '\n' {
			count++
		}
	}
	return count
}

func CountCodeLines(content []byte) int {
	lines := strings.Split(string(content), "\n")
	count := 0
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" && !isCommentLine(line) && !isEmptyLine(line) {
			count++
		}
	}
	return count
}

func CountCommentLines(content []byte) int {
	lines := strings.Split(string(content), "\n")
	count := 0
	inBlockComment := false

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)

		if inBlockComment {
			count++
			if strings.Contains(trimmed, "*/") {
				inBlockComment = false
				trimmed = strings.TrimPrefix(trimmed, "*/")
				trimmed = strings.TrimSpace(trimmed)
			}
		}

		if strings.HasPrefix(trimmed, "//") || strings.HasPrefix(trimmed, "#") {
			count++
			continue
		}

		if strings.HasPrefix(trimmed, "/*") {
			count++
			inBlockComment = true
			if strings.Contains(trimmed, "*/") {
				inBlockComment = false
			}
			continue
		}
	}

	return count
}

func isCommentLine(line string) bool {
	trimmed := strings.TrimSpace(line)
	return strings.HasPrefix(trimmed, "//") ||
		strings.HasPrefix(trimmed, "#") ||
		strings.HasPrefix(trimmed, "/*") ||
		strings.HasPrefix(trimmed, "*") ||
		strings.HasPrefix(trimmed, "<!--") ||
		strings.HasPrefix(trimmed, "--")
}

func isEmptyLine(line string) bool {
	return strings.TrimSpace(line) == ""
}

func GenerateDirectoryTree(rootPath string, prefix string) string {
	var sb strings.Builder

	files, err := os.ReadDir(rootPath)
	if err != nil {
		return ""
	}

	var dirs []os.DirEntry
	var codeFiles []os.DirEntry

	for _, file := range files {
		if strings.HasPrefix(file.Name(), ".") {
			continue
		}
		if file.IsDir() && IsExcludedDir(file.Name(), nil) {
			continue
		}
		if file.IsDir() {
			dirs = append(dirs, file)
		} else {
			ext := filepath.Ext(file.Name())
			if IsCodeFile(ext, GetValidExtensions()) {
				codeFiles = append(codeFiles, file)
			}
		}
	}

	for i, dir := range dirs {
		isLast := i == len(dirs)+len(codeFiles)-1
		connector := "└── "
		if !isLast {
			connector = "├── "
		}

		sb.WriteString(prefix)
		sb.WriteString(connector)
		sb.WriteString(dir.Name())
		sb.WriteString("/\n")

		newPrefix := prefix
		if isLast {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}

		subtree := GenerateDirectoryTree(filepath.Join(rootPath, dir.Name()), newPrefix)
		sb.WriteString(subtree)
	}

	for i, file := range codeFiles {
		isLast := i == len(codeFiles)-1
		connector := "└── "
		if !isLast {
			connector = "├── "
		}

		sb.WriteString(prefix)
		sb.WriteString(connector)
		sb.WriteString(file.Name())
		sb.WriteString("\n")
	}

	return sb.String()
}

func GetFileExtension(filePath string) string {
	return strings.ToLower(filepath.Ext(filePath))
}

func GetFileNameWithoutExtension(filePath string) string {
	base := filepath.Base(filePath)
	ext := filepath.Ext(base)
	return strings.TrimSuffix(base, ext)
}

func GetLanguageByExtension(extension string) string {
	langMap := map[string]string{
		".go":    "go",
		".js":    "javascript",
		".jsx":   "javascript",
		".ts":    "typescript",
		".tsx":   "typescript",
		".java":  "java",
		".py":    "python",
		".pyw":   "python",
		".cpp":   "cpp",
		".cc":    "cpp",
		".cxx":   "cpp",
		".hpp":   "cpp",
		".hh":    "cpp",
		".hxx":   "cpp",
		".h":     "c",
		".c":     "c",
		".php":   "php",
		".phtml": "php",
	}
	return langMap[strings.ToLower(extension)]
}

func SanitizePath(path string) string {
	path = filepath.Clean(path)
	path = regexp.MustCompile(`[/\\]+`).ReplaceAllString(path, "/")
	return path
}

func GetProjectName(rootPath string) string {
	base := filepath.Base(rootPath)
	if base == "." || base == "" {
		cwd, _ := os.Getwd()
		base = filepath.Base(cwd)
	}
	return base
}

func CreateOutputPath(outputDir string, projectName string) string {
	if outputDir == "" {
		outputDir = "."
	}

	outputFile := filepath.Join(outputDir, "code.md")

	if _, err := os.Stat(outputFile); err == nil {
		base := filepath.Base(outputFile)
		ext := filepath.Ext(base)
		name := strings.TrimSuffix(base, ext)
		counter := 1
		for {
			newFileName := fmt.Sprintf("%s_%d%s", name, counter, ext)
			newPath := filepath.Join(outputDir, newFileName)
			if _, err := os.Stat(newPath); os.IsNotExist(err) {
				outputFile = newPath
				break
			}
			counter++
		}
	}

	return outputFile
}

func init() {
	_, _ = regexp.Compile("")
}
