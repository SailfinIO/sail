# Sail Framework

Sail is a modular, NestJS-inspired framework for Go. It is designed to be minimal in external dependencies and easy to extend.

## Project Structure

- **cmd/sail/**: Entry point for the framework demo or CLI.
- **internal/**: Core implementation details, including dependency injection, server, routing, and logging.
- **pkg/sail/**: Public API for users to bootstrap and interact with the framework.
- **pkg/middleware/**: Optional middleware implementations (e.g., CORS).

## Getting Started

1. **Initialize the module:**

   ```bash
   go mod init github.com/SailfinIO/sail
   ```

Breakdown of Key Components
cmd/
This directory contains the entry point for your framework. For instance, cmd/sail/main.go A CLI tool that initializes and runs the framework.
internal/core/
container.go: Implements a basic dependency injection (DI) container. This helps in managing service instances and wiring up dependencies between modules.
module.go: Defines how modules are registered, initialized, and run. Similar to NestJS modules, this can encapsulate a feature’s controllers, providers, and configuration.
internal/server/
http.go: Wraps Go’s built-in net/http package to create a simple HTTP server.
router.go: Handles route registration and middleware integration. This module could evolve to support more complex routing features.
internal/logger/
A simple logging abstraction that wraps the standard library’s log package (or any minimal logging solution) to provide consistency across the framework.
pkg/sail/
This is your public-facing API. It exposes the core functionalities of your framework (like bootstrapping an application, module registration, and configuration) so that developers can easily import and use Sail in their own projects.
pkg/middleware/
Here you can provide optional middleware (for instance, for CORS, security, or request logging). This is kept separate from the core to maintain modularity.
go.mod & README.md
The go.mod file initializes your project as a Go module. The README.md serves as the landing page for users to understand how to get started with Sail.
