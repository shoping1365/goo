package config

import (
	"log"
	"os"
	"strconv"
	"time"
)

/*
پکیج تنظیمات

این پکیج برای مدیریت تنظیمات برنامه و متغیرهای محیطی استفاده می‌شود.
*/

// Config ساختار تنظیمات برنامه
type Config struct {
	// تنظیمات سرور
	Server struct {
		Port         string
		ReadTimeout  time.Duration
		WriteTimeout time.Duration
	}

	// تنظیمات دیتابیس
	Database struct {
		Host     string
		Port     string
		User     string
		Password string
		Name     string
		SSLMode  string
	}

	// تنظیمات Redis
	Redis struct {
		Host     string
		Port     string
		Password string
		DB       int
	}

	// تنظیمات JWT
	JWT struct {
		SecretKey string
		Duration  time.Duration
	}
}

// Load تنظیمات را از متغیرهای محیطی بارگذاری می‌کند
func Load() *Config {
	config := &Config{}

	// تنظیمات سرور
	config.Server.Port = getEnvRequired("BACKEND_PORT")
	config.Server.ReadTimeout = getDurationEnv("SERVER_READ_TIMEOUT", 10*time.Second)
	config.Server.WriteTimeout = getDurationEnv("SERVER_WRITE_TIMEOUT", 10*time.Second)

	// تنظیمات دیتابیس
	config.Database.Host = getEnvRequired("DB_HOST")
	config.Database.Port = getEnvRequired("DB_PORT")
	config.Database.User = getEnvRequired("DB_USER")
	config.Database.Password = getEnvRequired("DB_PASSWORD")
	config.Database.Name = getEnvRequired("DB_NAME")
	config.Database.SSLMode = getEnvRequired("DB_SSLMODE")

	// تنظیمات Redis
	config.Redis.Host = getEnvRequired("REDIS_HOST")
	config.Redis.Port = getEnvRequired("REDIS_PORT")
	config.Redis.Password = getEnv("REDIS_PASSWORD", "") // Password می‌تواند خالی باشد
	config.Redis.DB = getIntEnv("REDIS_DB", 0)

	// تنظیمات JWT
	config.JWT.SecretKey = getEnvRequired("JWT_SECRET_KEY")
	config.JWT.Duration = getDurationEnv("JWT_DURATION", 24*time.Hour)

	return config
}

// getEnvRequired مقدار متغیر محیطی را برمی‌گرداند و اگر تنظیم نشده باشد، خطا می‌دهد
func getEnvRequired(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("❌ Error: Required environment variable %s is not set", key)
	}
	return value
}

// getEnv مقدار متغیر محیطی را با مقدار پیش‌فرض برمی‌گرداند
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getIntEnv مقدار عددی متغیر محیطی را با مقدار پیش‌فرض برمی‌گرداند
func getIntEnv(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

// getDurationEnv مقدار زمانی متغیر محیطی را با مقدار پیش‌فرض برمی‌گرداند
func getDurationEnv(key string, defaultValue time.Duration) time.Duration {
	if value := os.Getenv(key); value != "" {
		if duration, err := time.ParseDuration(value); err == nil {
			return duration
		}
	}
	return defaultValue
}

func GetDatabaseURL() string {
	url := os.Getenv("DATABASE_URL")
	if url == "" {
		log.Fatal("❌ Error: DATABASE_URL environment variable is not set")
	}
	return url
}
