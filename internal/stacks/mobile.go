package stacks

import (
	"github.com/bigbytes/forge/internal/templates"
)

// MobileTemplateURL is the Git repo URL for mobile stack template
const MobileTemplateURL = "https://github.com/cmontecinos/forge-mobile"

// mobileStack implements the Stack interface for mobile projects
type mobileStack struct{}

func init() {
	Register(&mobileStack{})
}

func (s *mobileStack) ID() string {
	return "mobile"
}

func (s *mobileStack) Name() string {
	return "Mobile (Expo + Go + Supabase)"
}

func (s *mobileStack) Description() string {
	return "Expo React Native + Go backend + Supabase"
}

func (s *mobileStack) TemplateConfig() templates.TemplateConfig {
	return templates.TemplateConfig{
		RepoURL: MobileTemplateURL,
		Branch:  "main",
	}
}
