package stacks

import (
	"github.com/bigbytes/forge/internal/templates"
)

// WebTemplateURL is the Git repo URL for web stack template
const WebTemplateURL = "https://github.com/bigbytes/forge-template-web"

// webStack implements the Stack interface for web projects
type webStack struct{}

func init() {
	Register(&webStack{})
}

func (s *webStack) ID() string {
	return "web"
}

func (s *webStack) Name() string {
	return "Web (Next.js + Go + Supabase)"
}

func (s *webStack) Description() string {
	return "Next.js App Router + Go backend + Supabase"
}

func (s *webStack) TemplateConfig() templates.TemplateConfig {
	return templates.TemplateConfig{
		RepoURL: WebTemplateURL,
		Branch:  "main",
	}
}
