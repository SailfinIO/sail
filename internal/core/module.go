package core

// Module defines the interface for modules.
type Module interface {
	// Init is called to initialize the module.
	Init() error
}

// ModuleRegistry manages a list of modules.
type ModuleRegistry struct {
	modules []Module
}

// NewModuleRegistry creates a new module registry.
func NewModuleRegistry() *ModuleRegistry {
	return &ModuleRegistry{
		modules: []Module{},
	}
}

// Register adds a module to the registry.
func (mr *ModuleRegistry) Register(module Module) {
	mr.modules = append(mr.modules, module)
}

// InitAll initializes all registered modules.
func (mr *ModuleRegistry) InitAll() error {
	for _, module := range mr.modules {
		if err := module.Init(); err != nil {
			return err
		}
	}
	return nil
}
