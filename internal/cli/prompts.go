package cli

import (
	"github.com/AlecAivazis/survey/v2"
)

// StackType represents the type of project stack
type StackType string

const (
	StackWeb    StackType = "web"
	StackMobile StackType = "mobile"
)

// Feature represents optional features that can be added to a project
type Feature string

const (
	FeatureAuth     Feature = "auth"
	FeatureDatabase Feature = "database"
	FeatureAPI      Feature = "api"
)

// stackOptions maps display names to StackType values
var stackOptions = []string{
	"Web (Next.js + Go + Supabase)",
	"Mobile (Expo + Go + Supabase)",
}

// stackMapping maps display names to StackType
var stackMapping = map[string]StackType{
	"Web (Next.js + Go + Supabase)":    StackWeb,
	"Mobile (Expo + Go + Supabase)": StackMobile,
}

// featureOptions for multi-select display
var featureOptions = []string{
	"Auth - Login/registro via backend",
	"Database - Conexión Go-Supabase",
	"API - Router, middlewares, handlers",
}

// featureMapping maps display names to Feature values
var featureMapping = map[string]Feature{
	"Auth - Login/registro via backend":    FeatureAuth,
	"Database - Conexión Go-Supabase":      FeatureDatabase,
	"API - Router, middlewares, handlers": FeatureAPI,
}

// PromptStackSelection prompts the user to select a stack type
func PromptStackSelection() (StackType, error) {
	var selected string
	prompt := &survey.Select{
		Message: "Select stack:",
		Options: stackOptions,
	}

	if err := survey.AskOne(prompt, &selected); err != nil {
		return "", err
	}

	return stackMapping[selected], nil
}

// PromptFeatureSelection prompts the user to select features
func PromptFeatureSelection() ([]Feature, error) {
	var selected []string
	prompt := &survey.MultiSelect{
		Message: "Select features (space to toggle, enter to confirm):",
		Options: featureOptions,
	}

	if err := survey.AskOne(prompt, &selected); err != nil {
		return nil, err
	}

	features := make([]Feature, 0, len(selected))
	for _, s := range selected {
		if f, ok := featureMapping[s]; ok {
			features = append(features, f)
		}
	}

	return features, nil
}

// ParseStackType converts a string to StackType
func ParseStackType(s string) (StackType, bool) {
	switch s {
	case "web":
		return StackWeb, true
	case "mobile":
		return StackMobile, true
	default:
		return "", false
	}
}

// ParseFeatures converts a slice of strings to Features
func ParseFeatures(strs []string) []Feature {
	features := make([]Feature, 0, len(strs))
	for _, s := range strs {
		switch s {
		case "auth":
			features = append(features, FeatureAuth)
		case "database":
			features = append(features, FeatureDatabase)
		case "api":
			features = append(features, FeatureAPI)
		}
	}
	return features
}

// FormatFeatures returns a human-readable list of features
func FormatFeatures(features []Feature) string {
	if len(features) == 0 {
		return "none"
	}
	result := ""
	for i, f := range features {
		if i > 0 {
			result += ", "
		}
		result += string(f)
	}
	return result
}

// StackDisplayName returns the display name for a stack type
func StackDisplayName(st StackType) string {
	switch st {
	case StackWeb:
		return "Web (Next.js + Go + Supabase)"
	case StackMobile:
		return "Mobile (Expo + Go + Supabase)"
	default:
		return string(st)
	}
}
