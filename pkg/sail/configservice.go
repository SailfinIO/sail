package sail

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// ConfigService provides configuration management similar in spirit to NestJS's ConfigService.
type ConfigService struct {
	overrides map[string]string // Manual overrides set via Set().
	envs      map[string]string // Values loaded from a .env file.
}

// NewConfigService initializes a new ConfigService instance and loads the .env file if it exists.
func NewConfigService() *ConfigService {
	cs := &ConfigService{
		overrides: make(map[string]string),
		envs:      make(map[string]string),
	}
	cs.loadEnvFile(".env")
	return cs
}

// loadEnvFile loads key-value pairs from a .env file into the envs map.
func (cs *ConfigService) loadEnvFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		// .env file not found; skip loading.
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		// Skip empty lines and comments.
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		// Optionally remove surrounding quotes.
		if len(value) >= 2 && ((value[0] == '"' && value[len(value)-1] == '"') ||
			(value[0] == '\'' && value[len(value)-1] == '\'')) {
			value = value[1 : len(value)-1]
		}
		cs.envs[key] = value
	}
}

// Set allows you to manually override a configuration value.
func (cs *ConfigService) Set(key, value string) {
	cs.overrides[key] = value
}

// Get retrieves a configuration value as a string.
// Precedence: manual override > OS environment > .env file > default.
func (cs *ConfigService) Get(key string, defaultVal ...string) string {
	if val, ok := cs.overrides[key]; ok {
		return val
	}
	if env := os.Getenv(key); env != "" {
		return env
	}
	if val, ok := cs.envs[key]; ok {
		return val
	}
	if len(defaultVal) > 0 {
		return defaultVal[0]
	}
	return ""
}

// GetInt retrieves a configuration value as an integer.
func (cs *ConfigService) GetInt(key string, defaultVal ...int) int {
	str := cs.Get(key)
	if str == "" && len(defaultVal) > 0 {
		return defaultVal[0]
	}
	if i, err := strconv.Atoi(str); err == nil {
		return i
	}
	if len(defaultVal) > 0 {
		return defaultVal[0]
	}
	return 0
}

// GetBool retrieves a configuration value as a boolean.
func (cs *ConfigService) GetBool(key string, defaultVal ...bool) bool {
	str := cs.Get(key)
	if str == "" && len(defaultVal) > 0 {
		return defaultVal[0]
	}
	if b, err := strconv.ParseBool(str); err == nil {
		return b
	}
	if len(defaultVal) > 0 {
		return defaultVal[0]
	}
	return false
}
