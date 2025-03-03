package core

// Module defines the basic interface for a module in Sail.
// Modules should implement OnModuleInit to initialize themselves.
type Module interface {
	OnModuleInit() error
}

// Bootstrapper defines an optional interface for modules that need
// additional bootstrapping after initialization.
type Bootstrapper interface {
	OnApplicationBootstrap() error
}

// ShutdownHook defines an optional interface for modules that need
// to perform cleanup on shutdown.
type ShutdownHook interface {
	OnApplicationShutdown() error
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

// InitAll initializes all registered modules by calling OnModuleInit.
// Then it calls OnApplicationBootstrap for modules that implement Bootstrapper.
func (mr *ModuleRegistry) InitAll() error {
	for _, module := range mr.modules {
		if err := module.OnModuleInit(); err != nil {
			return err
		}
	}
	for _, module := range mr.modules {
		if bootstrapper, ok := module.(Bootstrapper); ok {
			if err := bootstrapper.OnApplicationBootstrap(); err != nil {
				return err
			}
		}
	}
	return nil
}

// ShutdownAll calls OnApplicationShutdown for modules that implement ShutdownHook.
func (mr *ModuleRegistry) ShutdownAll() error {
	for _, module := range mr.modules {
		if shutdownHook, ok := module.(ShutdownHook); ok {
			if err := shutdownHook.OnApplicationShutdown(); err != nil {
				return err
			}
		}
	}
	return nil
}
