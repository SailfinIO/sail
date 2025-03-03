package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// main serves as the CLI entry point.
func main() {
	// Define subcommands.
	newCmd := flag.NewFlagSet("new", flag.ExitOnError)
	generateCmd := flag.NewFlagSet("generate", flag.ExitOnError)

	// For the "new" command.
	var appName string
	newCmd.StringVar(&appName, "name", "", "Name of the new application")

	// For the "generate" command.
	var componentType string
	var componentName string
	generateCmd.StringVar(&componentType, "type", "", "Type of component to generate (module, controller, service)")
	generateCmd.StringVar(&componentName, "name", "", "Name of the component")

	// Check that a subcommand has been provided.
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	// Parse the subcommand.
	switch os.Args[1] {
	case "new":
		newCmd.Parse(os.Args[2:])
		if appName == "" {
			fmt.Println("Error: Please provide the application name with -name flag.")
			newCmd.Usage()
			os.Exit(1)
		}
		if err := createNewApp(appName); err != nil {
			fmt.Println("Error creating app:", err)
			os.Exit(1)
		}
		fmt.Println("App", appName, "created successfully.")
	case "generate":
		generateCmd.Parse(os.Args[2:])
		if componentType == "" || componentName == "" {
			fmt.Println("Error: Please provide both -type and -name flags.")
			generateCmd.Usage()
			os.Exit(1)
		}
		if err := generateComponent(componentType, componentName); err != nil {
			fmt.Println("Error generating component:", err)
			os.Exit(1)
		}
		fmt.Printf("Component %q of type %q created successfully.\n", componentName, componentType)
	default:
		printUsage()
		os.Exit(1)
	}
}

// printUsage prints a simple CLI usage guide.
func printUsage() {
	fmt.Println(`Usage: sail <command> [options]
Commands:
  new       -name <appName>            Creates a new application scaffold.
  generate  -type <componentType> -name <componentName>
             Generates a new component (module, controller, service).`)
}

// createNewApp scaffolds a new application.
func createNewApp(appName string) error {
	// Create the application directory.
	if err := os.Mkdir(appName, 0755); err != nil {
		return err
	}

	// Generate a go.mod file using the appName as the module name.
	goModPath := filepath.Join(appName, "go.mod")
	goModContent := fmt.Sprintf("module %s\n\ngo 1.24.0\n\nrequire github.com/SailfinIO/sail v0.0.2\n", appName)
	if err := os.WriteFile(goModPath, []byte(goModContent), 0644); err != nil {
		return err
	}

	// Create the main.go file. Note: import path for app is now "<appName>/app"
	mainFilePath := filepath.Join(appName, "main.go")
	mainContent := `package main

import (
	"github.com/SailfinIO/sail/pkg/sail"
	"` + appName + `/app"
)

func main() {
	// Create a new Sail app instance.
	appInstance := sail.NewApp()

	// Register the AppModule.
	appInstance.RegisterModule(&app.AppModule{})

	// Run the application.
	appInstance.Run()
}
`
	if err := os.WriteFile(mainFilePath, []byte(mainContent), 0644); err != nil {
		return err
	}

	// Create the "app" directory.
	appDir := filepath.Join(appName, "app")
	if err := os.Mkdir(appDir, 0755); err != nil {
		return err
	}

	// Create app.module.go.
	appModulePath := filepath.Join(appDir, "app.module.go")
	appModuleContent := `package app

// AppModule aggregates the application's components.
type AppModule struct{}

// OnModuleInit initializes the AppModule.
// This is where you can register controllers, services, or submodules.
func (m *AppModule) OnModuleInit() error {
	// TODO: Initialize your application's components.
	return nil
}
`
	if err := os.WriteFile(appModulePath, []byte(appModuleContent), 0644); err != nil {
		return err
	}

	// Create app.controller.go.
	appControllerPath := filepath.Join(appDir, "app.controller.go")
	appControllerContent := `package app

import (
	"net/http"

	"github.com/SailfinIO/sail/pkg/sail"
)

// AppController handles HTTP requests for the application.
type AppController struct {
	sail.BaseController
}

// RegisterRoutes registers the HTTP routes.
func (ac *AppController) RegisterRoutes(router *server.Router) {
	router.Handle("/", http.HandlerFunc(ac.index))
}

func (ac *AppController) index(w http.ResponseWriter, r *http.Request) {
	ac.WriteJSON(w, map[string]string{"message": "Welcome to your Sail application!"})
}
`
	if err := os.WriteFile(appControllerPath, []byte(appControllerContent), 0644); err != nil {
		return err
	}

	// Create app.service.go.
	appServicePath := filepath.Join(appDir, "app.service.go")
	appServiceContent := `package app

import (
	"github.com/SailfinIO/sail/pkg/sail"
)

// AppService encapsulates business logic for the application.
type AppService struct {
	sail.BaseService
}

// NewAppService creates a new instance of AppService.
func NewAppService(logger logger.Logger, config *sail.ConfigService) *AppService {
	return &AppService{
		BaseService: sail.NewBaseService(logger.WithContext("AppService"), config),
	}
}

// GetMessage returns a welcome message.
func (as *AppService) GetMessage() string {
	as.BaseService.Logger.Info("Retrieving welcome message")
	return "Hello from AppService!"
}
`
	if err := os.WriteFile(appServicePath, []byte(appServiceContent), 0644); err != nil {
		return err
	}

	return nil
}

// generateComponent scaffolds a new component based on its type.
func generateComponent(componentType, componentName string) error {
	switch componentType {
	case "module":
		// Create a directory for the module inside the modules directory.
		moduleDir := filepath.Join("modules", componentName)
		if err := os.MkdirAll(moduleDir, 0755); err != nil {
			return err
		}

		// Generate the module file.
		moduleFile := filepath.Join(moduleDir, strings.ToLower(componentName)+".module.go")
		moduleContent := `package ` + strings.ToLower(componentName) + `

import "github.com/SailfinIO/sail"

// ` + componentName + `Module is a basic module implementation.
type ` + componentName + `Module struct{}

// OnModuleInit initializes the module.
func (m *` + componentName + `Module) OnModuleInit() error {
	// TODO: Implement module initialization.
	return nil
}
`
		if err := os.WriteFile(moduleFile, []byte(moduleContent), 0644); err != nil {
			return err
		}

		// Generate the controller file.
		controllerFile := filepath.Join(moduleDir, strings.ToLower(componentName)+".controller.go")
		controllerContent := `package ` + strings.ToLower(componentName) + `

import (
	"net/http"
	"github.com/SailfinIO/sail"
)

// ` + componentName + `Controller handles HTTP requests for the ` + componentName + ` module.
type ` + componentName + `Controller struct {
	sail.BaseController
}

// RegisterRoutes registers the HTTP routes.
func (c *` + componentName + `Controller) RegisterRoutes(router *server.Router) {
	router.Handle("/` + strings.ToLower(componentName) + `", http.HandlerFunc(c.handle))
}

func (c *` + componentName + `Controller) handle(w http.ResponseWriter, r *http.Request) {
	c.WriteJSON(w, map[string]string{"message": "` + componentName + ` route"})
}
`
		if err := os.WriteFile(controllerFile, []byte(controllerContent), 0644); err != nil {
			return err
		}

		// Generate the service file.
		serviceFile := filepath.Join(moduleDir, strings.ToLower(componentName)+".service.go")
		serviceContent := `package ` + strings.ToLower(componentName) + `

import (
	"github.com/SailfinIO/sail"
)

// ` + componentName + `Service encapsulates business logic for the ` + componentName + ` module.
type ` + componentName + `Service struct {
	sail.BaseService
}

// New` + componentName + `Service creates a new instance of ` + componentName + `Service.
func New` + componentName + `Service(logger logger.Logger, config *sail.ConfigService) *` + componentName + `Service {
	return &` + componentName + `Service{
		BaseService: sail.NewBaseService(logger.WithContext("` + componentName + `Service"), config),
	}
}
`
		if err := os.WriteFile(serviceFile, []byte(serviceContent), 0644); err != nil {
			return err
		}

		return nil

	case "controller":
		// For a standalone controller, create it under controllers directory.
		controllerDir := "controllers"
		if err := os.MkdirAll(controllerDir, 0755); err != nil {
			return err
		}
		fileName := filepath.Join(controllerDir, componentName+".go")
		content := `package controllers

import (
	"net/http"
	"github.com/SailfinIO/sail"
)

// ` + componentName + `Controller is a basic controller implementation.
type ` + componentName + `Controller struct {
	sail.BaseController
}

// RegisterRoutes registers HTTP routes.
func (c *` + componentName + `Controller) RegisterRoutes(router *server.Router) {
	router.Handle("/` + strings.ToLower(componentName) + `", http.HandlerFunc(c.handle))
}

func (c *` + componentName + `Controller) handle(w http.ResponseWriter, r *http.Request) {
	c.WriteJSON(w, map[string]string{"message": "` + componentName + ` controller route"})
}
`
		if err := os.WriteFile(fileName, []byte(content), 0644); err != nil {
			return err
		}
		return nil

	case "service":
		// For a standalone service, create it under services directory.
		serviceDir := "services"
		if err := os.MkdirAll(serviceDir, 0755); err != nil {
			return err
		}
		fileName := filepath.Join(serviceDir, componentName+".go")
		content := `package services

import (
	"github.com/SailfinIO/sail"
)

// ` + componentName + `Service is a basic service implementation.
type ` + componentName + `Service struct {
	sail.BaseService
}

// New` + componentName + `Service creates a new instance of ` + componentName + `Service.
func New` + componentName + `Service(logger logger.Logger, config *sail.ConfigService) *` + componentName + `Service {
	return &` + componentName + `Service{
		BaseService: sail.NewBaseService(logger.WithContext("` + componentName + `Service"), config),
	}
}
`
		if err := os.WriteFile(fileName, []byte(content), 0644); err != nil {
			return err
		}
		return nil

	default:
		return fmt.Errorf("unknown component type: %s", componentType)
	}
}
