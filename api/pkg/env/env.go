package env

import (
	"os"
)

func GetBoolEnv(key string, defaultVal bool) bool {
	if val, exists := os.LookupEnv(key); exists {
		if val == "true" || val == "1" {
			return true
		}
		return false
	}
	return defaultVal
}

func GetStringEnv(key string, defaultVal string) string {
	if val, exists := os.LookupEnv(key); exists {
		return val
	}
	return defaultVal
}
