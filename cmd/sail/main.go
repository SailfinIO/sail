package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/SailfinIO/sail/cmd/sail/templates"
	"github.com/SailfinIO/sail/cmd/sail/version"
)

// Override the default version value.
// go build -ldflags "-X 'github.com/SailfinIO/sail/cmd/sail/cli/version.Version=v0.0.5'" -o sail ./cmd/sail/cli

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
		// For standalone generation, pass an empty baseDir so that the default directories are used.
		if err := generateComponent(componentType, componentName, ""); err != nil {
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
  new       -name <appName>            
             Creates a new application scaffold.
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
	goModContent := fmt.Sprintf(`module %s

go 1.24.0

require github.com/SailfinIO/sail %s
`, appName, version.Version)
	if err := os.WriteFile(goModPath, []byte(goModContent), 0644); err != nil {
		return err
	}

	// Create the main.go file.
	mainFilePath := filepath.Join(appName, "main.go")
	mainContent := `package main

import (
	"github.com/SailfinIO/sail/pkg/sail"
	"` + appName + `/app"
)

func main() {
	// Create a new Sail app instance.
	appInstance := sail.NewApp()

	// Create the AppModule.
	appModule := &app.AppModule{}

	// Register the AppModule.
	appInstance.RegisterModule(appModule)

	// Run the application.
	appInstance.Run()
}
`
	if err := os.WriteFile(mainFilePath, []byte(mainContent), 0644); err != nil {
		return err
	}

	// Create the "app" directory for the base application components.
	appDir := filepath.Join(appName, "app")
	if err := os.Mkdir(appDir, 0755); err != nil {
		return err
	}

	// Instead of manually writing the files, call the component generators with the base directory.
	// Here, we generate the module, controller, and service for the base app.
	if err := generateComponent("module", "App", appDir); err != nil {
		return err
	}
	if err := generateComponent("controller", "App", appDir); err != nil {
		return err
	}
	if err := generateComponent("service", "App", appDir); err != nil {
		return err
	}

	return nil
}

// generateComponent scaffolds a new component based on its type.
// The baseDir parameter allows the caller to specify a destination directory.
// If baseDir is an empty string, default directories will be used.
func generateComponent(componentType, componentName, baseDir string) error {
	switch componentType {
	case "module":
		return generateModule(componentName, baseDir)
	case "controller":
		return generateController(componentName, baseDir)
	case "service":
		return generateService(componentName, baseDir)
	default:
		return fmt.Errorf("unknown component type: %s", componentType)
	}
}

// generateFile creates a directory (if needed) and writes the file content.
func generateFile(dir, fileName, content string) error {
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	fullPath := filepath.Join(dir, fileName)
	return os.WriteFile(fullPath, []byte(content), 0644)
}

// generateModule creates a module file. If baseDir is provided it is used; otherwise, default to "modules".
func generateModule(componentName, baseDir string) error {
	// Use the provided baseDir or default directory.
	var moduleDir string
	if baseDir != "" {
		moduleDir = baseDir
	} else {
		moduleDir = filepath.Join("modules", strings.ToLower(componentName))
	}

	data := templates.ComponentData{
		Name:      componentName,
		LowerName: strings.ToLower(componentName),
	}

	// Render module file content from template.
	moduleContent, err := templates.RenderTemplate(templates.ModuleTemplate, data)
	if err != nil {
		return err
	}

	// Render controller file content from template.
	controllerContent, err := templates.RenderTemplate(templates.ControllerTemplate, data)
	if err != nil {
		return err
	}

	// Render service file content from template.
	serviceContent, err := templates.RenderTemplate(templates.ServiceTemplate, data)
	if err != nil {
		return err
	}

	files := []struct {
		fileName string
		content  string
	}{
		{
			fileName: data.LowerName + ".module.go",
			content:  moduleContent,
		},
		{
			fileName: data.LowerName + ".controller.go",
			content:  controllerContent,
		},
		{
			fileName: data.LowerName + ".service.go",
			content:  serviceContent,
		},
	}

	for _, f := range files {
		if err := generateFile(moduleDir, f.fileName, f.content); err != nil {
			return err
		}
	}
	return nil
}

// generateController creates a controller file. If baseDir is provided it is used; otherwise, default to "controllers".
func generateController(componentName, baseDir string) error {
	var controllerDir string
	if baseDir != "" {
		controllerDir = baseDir
	} else {
		controllerDir = "controllers"
	}

	data := templates.ComponentData{
		Name:      componentName,
		LowerName: strings.ToLower(componentName),
	}

	content, err := templates.RenderTemplate(templates.ControllerTemplate, data)
	if err != nil {
		return err
	}

	fileName := data.LowerName + ".controller.go"
	return generateFile(controllerDir, fileName, content)
}

func generateService(componentName, baseDir string) error {
	var serviceDir string
	if baseDir != "" {
		serviceDir = baseDir
	} else {
		serviceDir = "services"
	}

	data := templates.ComponentData{
		Name:      componentName,
		LowerName: strings.ToLower(componentName),
	}

	content, err := templates.RenderTemplate(templates.ServiceTemplate, data)
	if err != nil {
		return err
	}

	fileName := data.LowerName + ".service.go"
	return generateFile(serviceDir, fileName, content)
}
