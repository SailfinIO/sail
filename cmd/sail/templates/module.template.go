package templates

var ModuleTemplate = `package {{.LowerName}}

import (
	"fmt"
	"github.com/SailfinIO/sail/pkg/sail"
)

// {{.Name}}Module aggregates the module components.
type {{.Name}}Module struct {
	Controller *{{.Name}}Controller
	Service    *{{.Name}}Service
	app        *sail.App
}

// SetApp sets the app instance for the module.
func (m *{{.Name}}Module) SetApp(app *sail.App) {
	m.app = app
}

// OnModuleInit initializes the module by instantiating its controller and service,
// then registers the controller's routes.
func (m *{{.Name}}Module) OnModuleInit() error {
	if m.app == nil {
		return fmt.Errorf("app instance not set")
	}
	m.Controller = &{{.Name}}Controller{}
	m.Service = New{{.Name}}Service(sail.NewLogger().WithContext("{{.Name}}Service"), sail.NewConfigService())
	m.Controller.RegisterRoutes(m.app.Router())
	return nil
}
`
