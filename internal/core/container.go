package core

import "sync"

// Container provides a simple, thread-safe dependency injection container.
type Container struct {
	providers map[string]interface{}
	mu        sync.RWMutex
}

// NewContainer returns a new instance of Container.
func NewContainer() *Container {
	return &Container{
		providers: make(map[string]interface{}),
	}
}

// Register adds a provider to the container.
func (c *Container) Register(name string, provider interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.providers[name] = provider
}

// Resolve retrieves a provider by name.
func (c *Container) Resolve(name string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	p, ok := c.providers[name]
	return p, ok
}

// MustResolve retrieves a provider by name or panics if not found.
func (c *Container) MustResolve(name string) interface{} {
	if p, ok := c.Resolve(name); ok {
		return p
	}
	panic("provider not found: " + name)
}
