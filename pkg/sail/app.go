package sail

import (
	"github.com/SailfinIO/sail/internal/core"
	"github.com/SailfinIO/sail/internal/logger"
	"github.com/SailfinIO/sail/internal/server"
)

// App is the main application structure.
type App struct {
	container      *core.Container
	moduleRegistry *core.ModuleRegistry
	router         *server.Router
	httpServer     *server.HTTPServer
	logger         logger.Logger
}

// NewApp creates a new instance of App.
func NewApp() *App {
	container := core.NewContainer()
	moduleRegistry := core.NewModuleRegistry()
	router := server.NewRouter()
	logg := logger.New()
	return &App{
		container:      container,
		moduleRegistry: moduleRegistry,
		router:         router,
		logger:         logg,
	}
}

// RegisterModule registers a module with the application.
func (a *App) RegisterModule(module core.Module) {
	a.moduleRegistry.Register(module)
}

// Run initializes all modules and starts the HTTP server.
func (a *App) Run() {
	// Initialize modules
	if err := a.moduleRegistry.InitAll(); err != nil {
		a.logger.Error("Failed to initialize modules: " + err.Error())
		return
	}

	// Create and start HTTP server (default port set to :8080)
	a.httpServer = server.NewHTTPServer(":8080", a.router)
	a.logger.Info("Starting server on :8080")
	if err := a.httpServer.Start(); err != nil {
		a.logger.Error("Server error: " + err.Error())
	}
}
