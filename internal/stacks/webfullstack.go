package stacks

import (
	"github.com/bigbytes/forge/internal/templates"
)

// WebFullstackTemplateURL is the Git repo URL for web-fullstack stack template
const WebFullstackTemplateURL = "https://github.com/cmontecinos/forge-web-fullstack"

// webFullstackStack implements the Stack interface for web fullstack projects
type webFullstackStack struct{}

func init() {
	Register(&webFullstackStack{})
}

func (s *webFullstackStack) ID() string {
	return "web-fullstack"
}

func (s *webFullstackStack) Name() string {
	return "Web Fullstack (Next.js + Supabase)"
}

func (s *webFullstackStack) Description() string {
	return "Next.js App Router + Supabase (sin backend Go)"
}

func (s *webFullstackStack) TemplateConfig() templates.TemplateConfig {
	return templates.TemplateConfig{
		RepoURL: WebFullstackTemplateURL,
		Branch:  "main",
	}
}
