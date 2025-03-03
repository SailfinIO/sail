package server

import (
	"net/http"
)

// Router provides minimal routing functionality.
type Router struct {
	mux *http.ServeMux
}

// NewRouter returns a new Router instance.
func NewRouter() *Router {
	return &Router{
		mux: http.NewServeMux(),
	}
}

// Handle registers a new route with a handler.
func (r *Router) Handle(pattern string, handler http.Handler) {
	r.mux.Handle(pattern, handler)
}

// ServeHTTP makes Router implement the http.Handler interface.
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.mux.ServeHTTP(w, req)
}
