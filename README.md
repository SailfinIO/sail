# Sail Framework

Sail is a modular, NestJS-inspired framework for Go. It provides a lightweight and flexible platform for building web applications with minimal external dependencies.

## Project Structure

- **cmd/sail/cli**: Entry point for the `sail` CLI.
- **internal/**: Core implementation details including dependency injection, module lifecycle, HTTP server, routing, and logging.
- **pkg/sail/**: Public API for bootstrapping and interacting with the framework.
- **pkg/middleware/**: Optional middleware implementations (e.g., CORS).

## Features

- **Dependency Injection:** A simple, thread-safe DI container.
- **Module Lifecycle Hooks:** Modules can implement initialization, bootstrap, and shutdown hooks.
- **Graceful Server Shutdown:** Uses context-based shutdown for clean termination.
- **Middleware Support:** Chain multiple middleware functions for flexible request handling.
- **Config & Logging:** Environment‑driven configuration and log level support.

## Getting Started

1. **Initialize the module:**

   ```bash
   go mod init github.com/SailfinIO/sail
   ```
