package sail

import (
	"context"
	"github.com/SailfinIO/sail/internal/core"
	"github.com/SailfinIO/sail/internal/logger"
	"github.com/SailfinIO/sail/internal/server"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// App is the main application structure.
type App struct {
	container      *core.Container
	moduleRegistry *core.ModuleRegistry
	router         *server.Router
	httpServer     *server.HTTPServer
	logger         logger.Logger
	configService  *ConfigService
}

// NewApp creates a new instance of App.
func NewApp() *App {
	container := core.NewContainer()
	moduleRegistry := core.NewModuleRegistry()
	router := server.NewRouter()
	logg := logger.New()
	configService := NewConfigService()
	return &App{
		container:      container,
		moduleRegistry: moduleRegistry,
		router:         router,
		logger:         logg,
		configService:  configService,
	}
}

// RegisterModule registers a module with the application.
func (a *App) RegisterModule(module core.Module) {
	a.moduleRegistry.Register(module)
}

// Use adds a middleware to the application's router.
func (a *App) Use(mw server.Middleware) {
	a.router.Use(mw)
}

// Run initializes all modules and starts the HTTP server.
// It also listens for interrupt signals to gracefully shut down.
func (a *App) Run() {
	// Initialize modules.
	if err := a.moduleRegistry.InitAll(); err != nil {
		a.logger.Error("Failed to initialize modules: " + err.Error())
		return
	}

	// Determine server port via ConfigService (defaulting to 8080).
	addr := ":" + a.configService.Get("PORT", "8080")
	a.httpServer = server.NewHTTPServer(addr, a.router)
	a.logger.Info("Starting server on " + addr)

	// Start the HTTP server in a separate goroutine.
	serverErrChan := make(chan error, 1)
	go func() {
		if err := a.httpServer.Start(); err != nil {
			serverErrChan <- err
		}
	}()

	// Listen for interrupt signals for graceful shutdown.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	select {
	case err := <-serverErrChan:
		a.logger.Error("Server error: " + err.Error())
	case sig := <-quit:
		a.logger.Info("Received signal: " + sig.String() + ", shutting down...")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := a.httpServer.Shutdown(ctx); err != nil {
			a.logger.Error("Error during shutdown: " + err.Error())
		}
	}

	// Optionally call shutdown hooks for modules.
	if err := a.moduleRegistry.ShutdownAll(); err != nil {
		a.logger.Error("Error during module shutdown: " + err.Error())
	}
}
