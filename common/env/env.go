package env

import (
	"os"
	"strconv"
)

func GetString(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func GetInt(key string, fallback int) int {
	if value, ok := os.LookupEnv(key); ok {
		_value, err := strconv.Atoi(value)

		if err != nil {
			return fallback
		}

		return _value
	}

	return fallback
}

func GetBool(key string, fallback bool) bool {
	if value, ok := os.LookupEnv(key); ok {
		_value, err := strconv.ParseBool(value)

		if err != nil {
			return fallback
		}

		return _value
	}

	return fallback
}
