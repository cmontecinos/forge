package templates

// TemplateConfig holds configuration for a template
type TemplateConfig struct {
	// RepoURL is the Git repository URL
	RepoURL string

	// Branch to clone (default: "main")
	Branch string

	// Subdirectory within the repo to use (optional, for monorepo templates)
	Subdirectory string
}
