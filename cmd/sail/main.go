package main

import (
	"github.com/SailfinIO/sail/pkg/sail"
)

func main() {
	// Create a new app instance
	app := sail.NewApp()

	// Optionally, register modules or middleware
	// app.RegisterModule(...)

	// Run the application
	app.Run()
}
