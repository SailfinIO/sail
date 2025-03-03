package sail

import "github.com/SailfinIO/sail/internal/server"

// Router is the public alias for server.Router.
type Router = server.Router

// Middleware is the public alias for server.Middleware.
type Middleware = server.Middleware

// NewHTTPServer creates a new HTTPServer instance.
var NewHTTPServer = server.NewHTTPServer

// NewRouter creates a new Router instance.
var NewRouter = server.NewRouter
