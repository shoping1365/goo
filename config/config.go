package config

import (
	"os"
	"strconv"
)

// RedisConfig holds Redis configuration
type RedisConfig struct {
	Host     string
	Port     int
	Password string
	DB       int
}

// Config holds all configuration
type Config struct {
	// ... existing fields ...
	Redis RedisConfig
}

// getEnv مقدار متغیر محیطی را با مقدار پیش‌فرض برمی‌گرداند
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getEnvAsInt مقدار عددی متغیر محیطی را با مقدار پیش‌فرض برمی‌گرداند
func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

// LoadConfig loads configuration from environment variables
func LoadConfig() *Config {
	return &Config{
		// ... existing fields ...
		Redis: RedisConfig{
			Host:     getEnv("REDIS_HOST", "localhost"),
			Port:     getEnvAsInt("REDIS_PORT", 6379),
			Password: getEnv("REDIS_PASSWORD", ""),
			DB:       getEnvAsInt("REDIS_DB", 0),
		},
	}
}
