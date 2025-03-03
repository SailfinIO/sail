package server

import (
	"net/http"
)

// HTTPServer wraps the built-in net/http server.
type HTTPServer struct {
	server *http.Server
}

// NewHTTPServer creates a new HTTPServer instance.
func NewHTTPServer(addr string, handler http.Handler) *HTTPServer {
	return &HTTPServer{
		server: &http.Server{
			Addr:    addr,
			Handler: handler,
		},
	}
}

// Start runs the HTTP server.
func (s *HTTPServer) Start() error {
	return s.server.ListenAndServe()
}

// Shutdown stops the HTTP server.
func (s *HTTPServer) Shutdown() error {
	return s.server.Close()
}
