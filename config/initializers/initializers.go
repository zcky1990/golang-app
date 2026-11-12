// Package initializers provides a Rails-like boot hook.
//
// In Rails every file in config/initializers is loaded once during boot.
// In Go we emulate that with an explicit registry: packages register
// initializers with Register and main calls Run once at startup, right
// after the environment is loaded and before the database is opened.
package initializers

import (
	"fmt"
	"sync"
)

// Initializer is a named function executed once during application boot.
type Initializer struct {
	Name string
	Run  func() error
}

var (
	mu    sync.Mutex
	items []Initializer
)

// Register adds an initializer to the boot sequence. It is safe to call
// from package init() functions.
func Register(name string, fn func() error) {
	mu.Lock()
	defer mu.Unlock()
	items = append(items, Initializer{Name: name, Run: fn})
}

// Run executes every registered initializer in registration order.
func Run() error {
	mu.Lock()
	pending := make([]Initializer, len(items))
	copy(pending, items)
	mu.Unlock()

	for _, init := range pending {
		if init.Run == nil {
			continue
		}
		if err := init.Run(); err != nil {
			return fmt.Errorf("initializer %q failed: %w", init.Name, err)
		}
	}
	return nil
}

// Names returns the registered initializer names in execution order.
func Names() []string {
	mu.Lock()
	defer mu.Unlock()
	names := make([]string, 0, len(items))
	for _, init := range items {
		names = append(names, init.Name)
	}
	return names
}
