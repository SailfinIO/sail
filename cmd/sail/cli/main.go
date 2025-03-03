package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
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
	// Create the app directory.
	if err := os.Mkdir(appName, 0755); err != nil {
		return err
	}

	// Create a basic main.go file.
	mainFile := filepath.Join(appName, "main.go")
	mainContent := `package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to ` + appName + `!")
}
`
	if err := os.WriteFile(mainFile, []byte(mainContent), 0644); err != nil {
		return err
	}

	// Create additional directories for modules, controllers, and services.
	dirs := []string{"modules", "controllers", "services"}
	for _, dir := range dirs {
		path := filepath.Join(appName, dir)
		if err := os.Mkdir(path, 0755); err != nil {
			return err
		}
	}

	return nil
}

// generateComponent scaffolds a new component based on its type.
func generateComponent(componentType, componentName string) error {
	var dir, content string

	switch componentType {
	case "module":
		dir = "modules"
		content = `package modules

// ` + componentName + ` module implementation.
`
	case "controller":
		dir = "controllers"
		content = `package controllers

// ` + componentName + ` controller implementation.
`
	case "service":
		dir = "services"
		content = `package services

// ` + componentName + ` service implementation.
`
	default:
		return fmt.Errorf("unknown component type: %s", componentType)
	}

	// Create the directory if it does not exist.
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.Mkdir(dir, 0755); err != nil {
			return err
		}
	}

	// Create the file for the component.
	fileName := filepath.Join(dir, componentName+".go")
	if err := os.WriteFile(fileName, []byte(content), 0644); err != nil {
		return err
	}

	return nil
}
