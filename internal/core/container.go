package core

// Container provides a simple dependency injection container.
type Container struct {
	providers map[string]interface{}
}

// NewContainer returns a new instance of Container.
func NewContainer() *Container {
	return &Container{
		providers: make(map[string]interface{}),
	}
}

// Register adds a provider to the container.
func (c *Container) Register(name string, provider interface{}) {
	c.providers[name] = provider
}

// Resolve retrieves a provider by name.
func (c *Container) Resolve(name string) interface{} {
	return c.providers[name]
}
