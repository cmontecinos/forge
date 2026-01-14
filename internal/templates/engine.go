package templates

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/go-git/go-git/v5"
)

// TemplateData contains variables for template substitution
type TemplateData struct {
	ProjectName string
	StackType   string
	Features    []string
}

// CloneTemplate clones a Git repository to the destination directory
func CloneTemplate(repoURL, destDir string) error {
	_, err := git.PlainClone(destDir, false, &git.CloneOptions{
		URL:   repoURL,
		Depth: 1, // Shallow clone for speed
	})
	if err != nil {
		return WrapTemplateError("clone", err)
	}
	return nil
}

// ProcessTemplates walks the directory and processes all files as templates
func ProcessTemplates(dir string, data TemplateData) error {
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories
		if info.IsDir() {
			// Skip .git directory
			if info.Name() == ".git" {
				return filepath.SkipDir
			}
			return nil
		}

		// Read file content
		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		// Skip binary files (files containing null bytes)
		if isBinaryFile(content) {
			return nil
		}

		// Skip files that don't contain template markers
		if !containsTemplateMarkers(content) {
			return nil
		}

		// Parse and execute template
		tmpl, err := template.New(filepath.Base(path)).Parse(string(content))
		if err != nil {
			// Skip files that fail template parsing (might have Go template-like syntax but not meant to be templates)
			return nil
		}

		var buf bytes.Buffer
		if err := tmpl.Execute(&buf, data); err != nil {
			// Skip files that fail template execution
			return nil
		}

		// Write processed content back
		if err := os.WriteFile(path, buf.Bytes(), info.Mode()); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return WrapTemplateError("process", err)
	}
	return nil
}

// CleanupGitDir removes the .git directory after cloning
func CleanupGitDir(dir string) error {
	gitDir := filepath.Join(dir, ".git")
	if err := os.RemoveAll(gitDir); err != nil {
		return WrapTemplateError("cleanup", err)
	}
	return nil
}

// isBinaryFile checks if content contains null bytes (indicating binary)
func isBinaryFile(content []byte) bool {
	return bytes.Contains(content, []byte{0})
}

// containsTemplateMarkers checks if content has Go template syntax
func containsTemplateMarkers(content []byte) bool {
	return strings.Contains(string(content), "{{") && strings.Contains(string(content), "}}")
}

// HasFeature checks if a feature is in the features list
func (d TemplateData) HasFeature(feature string) bool {
	for _, f := range d.Features {
		if f == feature {
			return true
		}
	}
	return false
}
