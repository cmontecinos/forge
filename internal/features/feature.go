package features

// Feature defines the interface for an optional project feature
type Feature interface {
	// ID returns the unique identifier for this feature (e.g., "auth", "database")
	ID() string

	// Name returns the display name (e.g., "Auth - Login/registro via backend")
	Name() string

	// Description returns a longer description of the feature
	Description() string

	// CompatibleStacks returns stack IDs this feature works with, or nil for all stacks
	CompatibleStacks() []string
}
