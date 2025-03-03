package sail

// Config holds the configuration for the application.
type Config struct {
	Port string
}

// LoadConfig loads and returns the application configuration.
// This is a placeholder for more robust configuration management.
func LoadConfig() Config {
	return Config{
		Port: "8080",
	}
}
