package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

var (
	outputDir string
	force     bool
)

// projectNameRegex validates project names:
// - Must start with a letter
// - Can contain alphanumeric, hyphens, underscores
// - Max 64 characters
var projectNameRegex = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_-]{0,63}$`)

var newCmd = &cobra.Command{
	Use:   "new <project-name>",
	Short: "Create a new project",
	Long: `Create a new project with your preferred stack.

The project name must:
- Start with a letter
- Contain only letters, numbers, hyphens, and underscores
- Be at most 64 characters

Examples:
  forge new my-app
  forge new my-app --stack web
  forge new my-app --stack mobile --features auth,database,api
  forge new my-app -o /path/to/output`,
	Args: cobra.ExactArgs(1),
	RunE: runNew,
}

func init() {
	rootCmd.AddCommand(newCmd)

	newCmd.Flags().StringVarP(&outputDir, "output", "o", "", "Output directory for the new project")
	newCmd.Flags().BoolVarP(&force, "force", "f", false, "Overwrite existing directory")
}

func runNew(cmd *cobra.Command, args []string) error {
	projectName := args[0]

	// Validate project name
	if err := validateProjectName(projectName); err != nil {
		return err
	}

	// Determine output directory
	targetDir, err := resolveOutputDir(projectName)
	if err != nil {
		return err
	}

	// Check if directory exists
	if err := checkDirectory(targetDir); err != nil {
		return err
	}

	// Print confirmation (placeholder for now)
	fmt.Printf("Creating project: %s\n", projectName)
	fmt.Printf("Output: %s\n", targetDir)

	return nil
}

func validateProjectName(name string) error {
	if name == "" {
		return fmt.Errorf("project name cannot be empty")
	}

	if len(name) > 64 {
		return fmt.Errorf("project name must be at most 64 characters (got %d)", len(name))
	}

	if !projectNameRegex.MatchString(name) {
		if !isLetter(rune(name[0])) {
			return fmt.Errorf("project name must start with a letter: %s", name)
		}
		return fmt.Errorf("project name can only contain letters, numbers, hyphens, and underscores: %s", name)
	}

	return nil
}

func isLetter(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}

func resolveOutputDir(projectName string) (string, error) {
	var dir string

	if outputDir != "" {
		dir = outputDir
	} else {
		// Use project name as directory in current working directory
		cwd, err := os.Getwd()
		if err != nil {
			return "", fmt.Errorf("failed to get current directory: %w", err)
		}
		dir = filepath.Join(cwd, projectName)
	}

	// Resolve to absolute path
	absDir, err := filepath.Abs(dir)
	if err != nil {
		return "", fmt.Errorf("failed to resolve path: %w", err)
	}

	return absDir, nil
}

func checkDirectory(dir string) error {
	info, err := os.Stat(dir)
	if os.IsNotExist(err) {
		// Directory doesn't exist, that's fine
		return nil
	}
	if err != nil {
		return fmt.Errorf("failed to check directory: %w", err)
	}

	if !info.IsDir() {
		return fmt.Errorf("path exists but is not a directory: %s", dir)
	}

	// Directory exists, check if it's empty
	entries, err := os.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("failed to read directory: %w", err)
	}

	if len(entries) > 0 {
		if force {
			fmt.Printf("Warning: directory exists and is not empty, will be overwritten: %s\n", dir)
			return nil
		}
		return fmt.Errorf("directory exists and is not empty: %s\nUse --force to overwrite", dir)
	}

	return nil
}

// Helper to check if a string contains only valid characters (used for error messages)
func hasInvalidChars(name string) bool {
	for _, r := range name {
		if !isLetter(r) && !isDigit(r) && r != '-' && r != '_' {
			return true
		}
	}
	return false
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

// GetInvalidChars returns the invalid characters in a project name
func GetInvalidChars(name string) string {
	var invalid []rune
	for _, r := range name {
		if !isLetter(r) && !isDigit(r) && r != '-' && r != '_' {
			invalid = append(invalid, r)
		}
	}
	return string(invalid)
}

// FormatInvalidChars returns a human-readable list of invalid characters
func FormatInvalidChars(name string) string {
	chars := GetInvalidChars(name)
	if chars == "" {
		return ""
	}
	var parts []string
	for _, r := range chars {
		parts = append(parts, fmt.Sprintf("'%c'", r))
	}
	return strings.Join(parts, ", ")
}
