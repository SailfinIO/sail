package templates

var ControllerTemplate = `package {{.LowerName}}

import (
	"net/http"
	"github.com/SailfinIO/sail/pkg/sail"
)

// {{.Name}}Controller handles HTTP requests.
type {{.Name}}Controller struct {
	sail.BaseController
}

// RegisterRoutes registers HTTP routes.
func (c *{{.Name}}Controller) RegisterRoutes(router *sail.Router) {
	router.Handle("/{{.LowerName}}", http.HandlerFunc(c.handle))
}

func (c *{{.Name}}Controller) handle(w http.ResponseWriter, r *http.Request) {
	c.WriteJSON(w, map[string]string{"message": "Welcome to your Sail application from {{.Name}}!"})
}
`
