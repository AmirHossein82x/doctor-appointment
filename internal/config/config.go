package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"sync"
)

// Config struct to hold environment variables
type Config struct {
	DB_HOST           string
	DB_PORT           string
	DB_USER           string
	DB_PASSWORD       string
	DB_NAME           string
	DB_SSLMODE        string
	REDIS_ADDR        string
	REDIS_DB          string
	ELASTIC_USERNAME  string
	ELASTIC_PASSWORD  string
	ELASTIC_ADDR      string
	ELASTIC_INDEX     string
	KAVENEGAR_API_KEY string
	KAVENEGAR_SENDER  string
}

// Singleton instance and mutex to make it thread-safe
var configInstance *Config
var once sync.Once

// LoadConfig returns the singleton instance of Config, ensuring it is only created once
func LoadConfig() *Config {
	once.Do(func() {
		// Load environment variables
		err := godotenv.Load()
		if err != nil {
			log.Println("Warning: No .env file found, using system environment variables")
		}

		// Initialize the config instance
		configInstance = &Config{
			DB_HOST:           getEnv("DB_HOST", "default"),
			DB_PORT:           getEnv("DB_PORT", "default"),
			DB_USER:           getEnv("DB_USER", "default"),
			DB_PASSWORD:       getEnv("DB_PASSWORD", "default"),
			DB_NAME:           getEnv("DB_NAME", "default"),
			DB_SSLMODE:        getEnv("DB_SSLMODE", "default"),
			REDIS_ADDR:        getEnv("REDIS_ADDR", "default"),
			REDIS_DB:          getEnv("REDIS_DB", "default"),
			ELASTIC_USERNAME:  getEnv("ELASTIC_USERNAME", "default"),
			ELASTIC_PASSWORD:  getEnv("ELASTIC_PASSWORD", "default"),
			ELASTIC_ADDR:      getEnv("ELASTIC_ADDR", "default"),
			ELASTIC_INDEX:     getEnv("ELASTIC_INDEX", "default"),
			KAVENEGAR_API_KEY: getEnv("KAVENEGAR_API_KEY", "default"),
			KAVENEGAR_SENDER:  getEnv("KAVENEGAR_SENDER", "default"),
		}
	})

	return configInstance
}

// getEnv retrieves an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
