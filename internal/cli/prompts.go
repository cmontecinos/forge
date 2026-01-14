package cli

import (
	"errors"

	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/terminal"
	"github.com/bigbytes/forge/internal/features"
	"github.com/bigbytes/forge/internal/stacks"
)

// StackType represents the type of project stack (maps to stack ID)
type StackType string

// Feature represents optional features that can be added to a project (feature ID)
type Feature = string

// PromptStackSelection prompts the user to select a stack type
func PromptStackSelection() (StackType, error) {
	// Build options dynamically from registry
	allStacks := stacks.All()
	options := make([]string, len(allStacks))
	for i, s := range allStacks {
		options[i] = s.Name()
	}

	var selected string
	prompt := &survey.Select{
		Message: "Select stack:",
		Options: options,
	}

	if err := survey.AskOne(prompt, &selected); err != nil {
		if errors.Is(err, terminal.InterruptErr) {
			return "", &UserAbortError{}
		}
		return "", err
	}

	// Find stack ID from selected name
	for _, s := range allStacks {
		if s.Name() == selected {
			return StackType(s.ID()), nil
		}
	}

	return "", nil
}

// PromptFeatureSelection prompts the user to select features
func PromptFeatureSelection() ([]Feature, error) {
	// Build options dynamically from registry
	allFeatures := features.All()
	options := make([]string, len(allFeatures))
	for i, f := range allFeatures {
		options[i] = f.Name()
	}

	var selected []string
	prompt := &survey.MultiSelect{
		Message: "Select features (space to toggle, enter to confirm):",
		Options: options,
	}

	if err := survey.AskOne(prompt, &selected); err != nil {
		if errors.Is(err, terminal.InterruptErr) {
			return nil, &UserAbortError{}
		}
		return nil, err
	}

	// Convert selected names to feature IDs
	result := make([]Feature, 0, len(selected))
	for _, name := range selected {
		for _, f := range allFeatures {
			if f.Name() == name {
				result = append(result, f.ID())
				break
			}
		}
	}

	return result, nil
}

// ParseStackType converts a string to StackType using the registry
func ParseStackType(s string) (StackType, bool) {
	_, ok := stacks.Get(s)
	if ok {
		return StackType(s), true
	}
	return "", false
}

// ParseFeatures converts a slice of strings to Features using the registry
func ParseFeatures(strs []string) []Feature {
	result := make([]Feature, 0, len(strs))
	for _, s := range strs {
		if _, ok := features.Get(s); ok {
			result = append(result, s)
		}
	}
	return result
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

// StackDisplayName returns the display name for a stack type using the registry
func StackDisplayName(st StackType) string {
	s, ok := stacks.Get(string(st))
	if ok {
		return s.Name()
	}
	return string(st)
}
