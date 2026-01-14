package templates

// Default template repository URLs
const (
	// WebTemplateURL is the Git repo URL for web stack template
	WebTemplateURL = "https://github.com/bigbytes/forge-template-web"

	// MobileTemplateURL is the Git repo URL for mobile stack template
	MobileTemplateURL = "https://github.com/bigbytes/forge-template-mobile"
)

// TemplateConfig holds configuration for a template
type TemplateConfig struct {
	// RepoURL is the Git repository URL
	RepoURL string

	// Branch to clone (default: "main")
	Branch string

	// Subdirectory within the repo to use (optional, for monorepo templates)
	Subdirectory string
}

// DefaultWebConfig returns the default configuration for web stack
func DefaultWebConfig() TemplateConfig {
	return TemplateConfig{
		RepoURL: WebTemplateURL,
		Branch:  "main",
	}
}

// DefaultMobileConfig returns the default configuration for mobile stack
func DefaultMobileConfig() TemplateConfig {
	return TemplateConfig{
		RepoURL: MobileTemplateURL,
		Branch:  "main",
	}
}

// GetConfigForStack returns the appropriate config for a stack type
func GetConfigForStack(stackType string) TemplateConfig {
	switch stackType {
	case "web":
		return DefaultWebConfig()
	case "mobile":
		return DefaultMobileConfig()
	default:
		return DefaultWebConfig()
	}
}
