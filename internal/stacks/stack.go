package stacks

import (
	"github.com/bigbytes/forge/internal/templates"
)

// Stack defines the interface for a project stack
type Stack interface {
	// ID returns the unique identifier for this stack (e.g., "web", "mobile")
	ID() string

	// Name returns the display name (e.g., "Web (Next.js + Go + Supabase)")
	Name() string

	// Description returns a short description of the stack
	Description() string

	// TemplateConfig returns the template configuration for this stack
	TemplateConfig() templates.TemplateConfig
}
