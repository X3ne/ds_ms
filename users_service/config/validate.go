package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func ValidateEnvs() map[string]string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	logLevels := map[string]bool{
		"DEBUG":   true,
		"INFO":    true,
		"WARNING": true,
		"ERROR":   true,
		"FATAL":   true,
	}

	envs := map[string]string{
		"LOG_LEVEL": os.Getenv("LOG_LEVEL"),
		"HOST":      os.Getenv("HOST"),
		"PORT":      os.Getenv("PORT"),
		"DB_PATH":   os.Getenv("DB_PATH"),
	}

	for key, value := range envs {
		if value == "" {
			log.Fatalf("Missing %s environment variable", key)
		}
	}

	if _, ok := logLevels[envs["LOG_LEVEL"]]; !ok {
		log.Fatalf("Invalid LOG_LEVEL environment variable. Valid values are: DEBUG, INFO, WARNING, ERROR, FATAL")
	}

	return envs
}
