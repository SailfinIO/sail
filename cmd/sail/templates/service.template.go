package templates

var ServiceTemplate = `package {{.LowerName}}

import (
	"github.com/SailfinIO/sail/pkg/sail"
)

// {{.Name}}Service encapsulates business logic.
type {{.Name}}Service struct {
	sail.BaseService
}

// New{{.Name}}Service creates a new instance of {{.Name}}Service.
func New{{.Name}}Service(logger sail.Logger, config *sail.ConfigService) *{{.Name}}Service {
	return &{{.Name}}Service{
		BaseService: sail.NewBaseService(logger.WithContext("{{.Name}}Service"), config),
	}
}

// GetMessage returns a welcome message.
func (as *{{.Name}}Service) GetMessage() string {
	as.Logger.Info("Retrieving welcome message")
	return "Hello from {{.Name}}Service!"
}
`
