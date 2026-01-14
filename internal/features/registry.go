package features

import (
	"sort"
	"sync"
)

var (
	registry = make(map[string]Feature)
	mu       sync.RWMutex
)

// Register adds a feature to the global registry
func Register(feature Feature) {
	mu.Lock()
	defer mu.Unlock()
	registry[feature.ID()] = feature
}

// Get retrieves a feature by ID from the registry
func Get(id string) (Feature, bool) {
	mu.RLock()
	defer mu.RUnlock()
	feature, ok := registry[id]
	return feature, ok
}

// All returns all registered features, sorted by ID for consistent ordering
func All() []Feature {
	mu.RLock()
	defer mu.RUnlock()

	features := make([]Feature, 0, len(registry))
	for _, feature := range registry {
		features = append(features, feature)
	}

	// Sort by ID for consistent order
	sort.Slice(features, func(i, j int) bool {
		return features[i].ID() < features[j].ID()
	})

	return features
}

// IDs returns all registered feature IDs, sorted
func IDs() []string {
	mu.RLock()
	defer mu.RUnlock()

	ids := make([]string, 0, len(registry))
	for id := range registry {
		ids = append(ids, id)
	}

	sort.Strings(ids)
	return ids
}

// ForStack returns features compatible with the given stack ID
func ForStack(stackID string) []Feature {
	mu.RLock()
	defer mu.RUnlock()

	features := make([]Feature, 0, len(registry))
	for _, feature := range registry {
		compatible := feature.CompatibleStacks()
		if compatible == nil {
			// nil means compatible with all stacks
			features = append(features, feature)
		} else {
			for _, sid := range compatible {
				if sid == stackID {
					features = append(features, feature)
					break
				}
			}
		}
	}

	// Sort by ID for consistent order
	sort.Slice(features, func(i, j int) bool {
		return features[i].ID() < features[j].ID()
	})

	return features
}
