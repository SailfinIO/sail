package sail

import "github.com/SailfinIO/sail/internal/logger"

// BaseService provides common shared dependencies for services,
// such as logging and configuration.
type BaseService struct {
	Logger logger.Logger  // Logger for logging messages.
	Config *ConfigService // ConfigService for accessing configuration values.
}

// NewBaseService returns a new instance of BaseService initialized
// with the provided logger and config service.
func NewBaseService(logger logger.Logger, config *ConfigService) BaseService {
	return BaseService{
		Logger: logger,
		Config: config,
	}
}

// Service is a marker interface for business logic components.
type Service interface{}
