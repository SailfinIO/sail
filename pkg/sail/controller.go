package sail

import (
	"encoding/json"
	"net/http"

	"github.com/SailfinIO/sail/internal/server"
)

// Controller defines the contract for HTTP controllers.
// Controllers should implement RegisterRoutes to bind their endpoints to the router.
type Controller interface {
	RegisterRoutes(router *server.Router)
}

// BaseController can be embedded by controllers to reuse common functionality.
type BaseController struct{}

// WriteJSON is a helper to write data as JSON with appropriate headers.
func (bc *BaseController) WriteJSON(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(data)
}

// ReadJSON is a helper to decode JSON from the request body.
func (bc *BaseController) ReadJSON(r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}
