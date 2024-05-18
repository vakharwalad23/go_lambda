package config

import (
	"os"
)

type Config struct {
	JWTSecret string
}

func NewConfig() *Config {
	return &Config{
		JWTSecret: getEnv("JWT_SECRET", ""),
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
