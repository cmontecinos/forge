package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/bigbytes/forge/internal/stacks"
	"github.com/bigbytes/forge/internal/templates"
	"github.com/spf13/cobra"
)

var (
	outputDir    string
	force        bool
	stackFlag    string
	featuresFlag string
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
- Be at most 64 characters`,
	Example: `  # Interactive mode - prompts for stack and features
  forge new my-app

  # Specify stack, prompt for features
  forge new my-app --stack web

  # Fully non-interactive
  forge new my-app --stack mobile --features auth,database,api

  # Custom output directory
  forge new my-app -o /path/to/projects

  # Overwrite existing directory
  forge new my-app --force`,
	Args: cobra.ExactArgs(1),
	RunE: runNew,
}

func init() {
	rootCmd.AddCommand(newCmd)

	newCmd.Flags().StringVarP(&outputDir, "output", "o", "", "Output directory for the new project")
	newCmd.Flags().BoolVarP(&force, "force", "f", false, "Overwrite existing directory")
	newCmd.Flags().StringVarP(&stackFlag, "stack", "s", "", "Stack type (web or mobile)")
	newCmd.Flags().StringVar(&featuresFlag, "features", "", "Features to include (comma-separated: auth,database,api)")
}

func runNew(cmd *cobra.Command, args []string) error {
	startTime := time.Now()
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

	// Get stack selection (flag or prompt)
	stack, err := getStackSelection(cmd)
	if err != nil {
		if IsUserAbort(err) {
			fmt.Println("\nAborted.")
			return nil
		}
		return err
	}

	// Get features selection (flag or prompt)
	features, err := getFeaturesSelection(cmd)
	if err != nil {
		if IsUserAbort(err) {
			fmt.Println("\nAborted.")
			return nil
		}
		return err
	}

	// Print summary
	PrintSummary(projectName, stack, features, targetDir)

	// If force flag is set and directory exists, remove it first
	if force {
		if _, err := os.Stat(targetDir); err == nil {
			Info("Removing existing directory...")
			if err := os.RemoveAll(targetDir); err != nil {
				return fmt.Errorf("failed to remove existing directory: %w", err)
			}
		}
	}

	// Get template configuration for selected stack from registry
	s, _ := stacks.Get(string(stack))
	config := s.TemplateConfig()

	// Convert features to string slice for TemplateData
	featureStrings := make([]string, len(features))
	for i, f := range features {
		featureStrings[i] = string(f)
	}

	// Prepare template data
	data := templates.TemplateData{
		ProjectName: projectName,
		StackType:   string(stack),
		Features:    featureStrings,
	}

	// Create project using template engine
	Info("Cloning template...")
	if err := templates.CloneTemplate(config.RepoURL, targetDir); err != nil {
		Error("Failed to clone template")
		return err
	}

	Info("Processing templates...")
	if err := templates.ProcessTemplates(targetDir, data); err != nil {
		Error("Failed to process templates")
		return err
	}

	Info("Cleaning up...")
	if err := templates.CleanupGitDir(targetDir); err != nil {
		// Non-fatal but worth mentioning
		Warn("Could not remove .git directory")
	}

	duration := time.Since(startTime).Seconds()
	Success("Project created successfully!")
	PrintNextSteps(projectName, stack, duration)

	return nil
}

func getStackSelection(cmd *cobra.Command) (StackType, error) {
	// If flag provided, use it
	if cmd.Flags().Changed("stack") {
		st, ok := ParseStackType(stackFlag)
		if !ok {
			// Build list of valid stack IDs for error message
			ids := stacks.IDs()
			return "", fmt.Errorf("invalid stack type: %s (must be one of: %s)", stackFlag, strings.Join(ids, ", "))
		}
		return st, nil
	}

	// Otherwise, prompt
	return PromptStackSelection()
}

func getFeaturesSelection(cmd *cobra.Command) ([]Feature, error) {
	// If flag was explicitly set (even to empty), use it
	if cmd.Flags().Changed("features") {
		if featuresFlag == "" {
			return []Feature{}, nil
		}
		featureStrs := strings.Split(featuresFlag, ",")
		// Trim whitespace and filter empty strings
		var cleanStrs []string
		for _, f := range featureStrs {
			f = strings.TrimSpace(f)
			if f != "" {
				cleanStrs = append(cleanStrs, f)
			}
		}
		// Validate all features are recognized
		for _, f := range cleanStrs {
			if f != "auth" && f != "database" && f != "api" {
				return nil, fmt.Errorf("invalid feature: %s (must be auth, database, or api)", f)
			}
		}
		return ParseFeatures(cleanStrs), nil
	}

	// Otherwise, prompt
	return PromptFeatureSelection()
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
