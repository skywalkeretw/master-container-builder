package pkg

import (
	"os"
	"strconv"
)

func GetEnvSting(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}

func GetEnvInt(key string, defaultValue int) int {
	envValue, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	value, err := strconv.Atoi(envValue)

	if err != nil {
		return defaultValue
	}

	return value
}
