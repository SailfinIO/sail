package logger

import "log"

// Logger defines a simple logging interface.
type Logger interface {
	Info(msg string)
	Error(msg string)
}

// DefaultLogger is a basic logger using the standard log package.
type DefaultLogger struct{}

// Info logs informational messages.
func (l *DefaultLogger) Info(msg string) {
	log.Println("INFO:", msg)
}

// Error logs error messages.
func (l *DefaultLogger) Error(msg string) {
	log.Println("ERROR:", msg)
}

// New returns a new instance of the default logger.
func New() Logger {
	return &DefaultLogger{}
}
