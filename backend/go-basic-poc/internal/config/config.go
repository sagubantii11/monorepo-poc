package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port         string
	ClientID     string
	ClientSecret string
	DBHost       string
	DBPort       string
	DBUser       string
	DBPassword   string
	DBName       string
	DBSSLMode    string
	// Add other configuration fields here
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using environment variables")
	}

	cfg := &Config{
		Port:         os.Getenv("PORT"),
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		DBHost:       os.Getenv("DB_HOST"),
		DBPort:       os.Getenv("DB_PORT"),
		DBUser:       os.Getenv("DB_USER"),
		DBPassword:   os.Getenv("DB_PASSWORD"),
		DBName:       os.Getenv("DB_NAME"),
		DBSSLMode:    os.Getenv("DB_SSLMODE"),
		// Load other configuration fields from environment variables
	}

	if cfg.Port == "" {
		cfg.Port = "8080" // Default port
	}

	// Set default values if needed.
	if cfg.DBPort == "" {
		cfg.DBPort = "5432" // Default PostgreSQL port
	}

	if cfg.DBSSLMode == "" {
		cfg.DBSSLMode = "disable" // Default ssl mode
	}

	return cfg, nil
}
