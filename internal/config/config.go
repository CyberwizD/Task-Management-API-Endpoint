package config

import (
	"os"

	"github.com/gin-gonic/gin"
)

// Config holds application configuration
type Config struct {
	Port    string
	GinMode string
	Env     string
}

// Load loads configuration from environment variables with defaults
func Load() *Config {
	return &Config{
		Port:    getEnv("PORT", "8080"),
		GinMode: getGinMode(),
		Env:     getEnv("APP_ENV", "development"),
	}
}

// getEnv gets environment variable with fallback to default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getGinMode determines Gin mode based on environment
func getGinMode() string {
	env := getEnv("APP_ENV", "development")
	
	switch env {
	case "production":
		return gin.ReleaseMode
	case "test":
		return gin.TestMode
	default:
		return gin.DebugMode
	}
}

// IsDevelopment checks if the application is running in development mode
func (c *Config) IsDevelopment() bool {
	return c.Env == "development"
}

// IsProduction checks if the application is running in production mode
func (c *Config) IsProduction() bool {
	return c.Env == "production"
}
