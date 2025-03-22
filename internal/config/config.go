package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

// Config struct to hold environment variables
type Config struct {
	DB_HOST     string
	DB_PORT     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
	DB_SSLMODE  string
	REDIS_ADDR  string
	REDIS_DB    string
}

// LoadConfig reads the .env file and maps it to the Config struct
func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found, using system environment variables")
	}

	return &Config{
		DB_HOST:     getEnv("DB_HOST", "default"),
		DB_PORT:     getEnv("DB_PORT", "default"),
		DB_USER:     getEnv("DB_USER", "default"),
		DB_PASSWORD: getEnv("DB_PASSWORD", "default"),
		DB_NAME:     getEnv("DB_NAME", "default"),
		DB_SSLMODE:  getEnv("DB_SSLMODE", "default"),
		REDIS_ADDR:  getEnv("REDIS_ADDR", "default"),
		REDIS_DB:    getEnv("REDIS_DB", "default"),
	}
}

// getEnv retrieves an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
