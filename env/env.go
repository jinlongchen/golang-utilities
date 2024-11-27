package env

import (
	"os"
)

// GetEnv retrieves the value of the environment variable named by the key.
// If the variable is not present, it returns the default value.
func GetEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
