package retrieve_env

import "os"

// RetrieveEnv retrieves an environment variable or returns a fallback value
func RetrieveEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	return value
}
