package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	ServerHost string
	ServerPort string
	GinMode    string
}

var AppConfig *Config

func LoadConfig() {
	var err error = godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	AppConfig = &Config{
		ServerHost: getEnvValue("SERVER_HOST", "0.0.0.0"),
		ServerPort: getEnvValue("SERVER_PORT", "8080"),
		GinMode:    getEnvValue("GIN_MODE", "debug"),
	}
}

func getEnvValue(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}
