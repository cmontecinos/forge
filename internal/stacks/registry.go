package stacks

import (
	"sort"
	"sync"
)

var (
	registry = make(map[string]Stack)
	mu       sync.RWMutex
)

// Register adds a stack to the global registry
func Register(stack Stack) {
	mu.Lock()
	defer mu.Unlock()
	registry[stack.ID()] = stack
}

// Get retrieves a stack by ID from the registry
func Get(id string) (Stack, bool) {
	mu.RLock()
	defer mu.RUnlock()
	stack, ok := registry[id]
	return stack, ok
}

// All returns all registered stacks, sorted by ID for consistent ordering
func All() []Stack {
	mu.RLock()
	defer mu.RUnlock()

	stacks := make([]Stack, 0, len(registry))
	for _, stack := range registry {
		stacks = append(stacks, stack)
	}

	// Sort by ID for consistent order
	sort.Slice(stacks, func(i, j int) bool {
		return stacks[i].ID() < stacks[j].ID()
	})

	return stacks
}

// IDs returns all registered stack IDs, sorted
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
