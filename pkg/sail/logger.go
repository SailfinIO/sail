package sail

import "github.com/SailfinIO/sail/internal/logger"

// Logger is the public alias for logger.Logger.
type Logger = logger.Logger

// NewLogger is a convenience function to create a new logger.
var NewLogger = logger.New
