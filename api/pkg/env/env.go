package env

import (
	"fmt"
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

func GetIntEnv(key string, defaultVal int) int {
	if val, exists := os.LookupEnv(key); exists {
		var intVal int
		_, err := fmt.Sscanf(val, "%d", &intVal)
		if err == nil {
			return intVal
		}
	}
	return defaultVal
}
