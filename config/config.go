package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	APIPort int

	PostgresDB       string
	PostgresUser     string
	PostgresPassword string
	PostgresHost     string
	PostgresPort     int
}

func LoadFromEnv() *Config {
	_ = godotenv.Load()

	cfg := &Config{
		APIPort:            getEnvAsInt("API_PORT", 8000),
		PostgresDB:          getEnv("POSTGRES_DB", "url_shortener"),
		PostgresUser:        getEnv("POSTGRES_USER", "postgres"),
		PostgresPassword:    getEnv("POSTGRES_PASSWORD", "postgres"),
		PostgresHost:        getEnv("POSTGRES_HOST", "postgres"),
		PostgresPort:        getEnvAsInt("POSTGRES_PORT", 5432),
	}

	return cfg
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return defaultValue
	}
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		fmt.Printf("Warning: invalid integer value for %s: %s, using default: %d\n", key, valueStr, defaultValue)
		return defaultValue
	}
	return value
}

