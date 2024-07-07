package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
}

// LoadConfig loads configuration settings from environment variables and .env file.
// It returns a pointer to a Config struct and an error if any.
func LoadConfig() (*Config, error) {
	// Load environment variables from .env file.
	// If .env file does not exist or cannot be loaded, godotenv.Load returns an error.
	err := godotenv.Load(".env")
	if err != nil {
		// Log the error and terminate the program.
		// The %v placeholder in the log message will be replaced with the error value.
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Initialize a Config struct with values from environment variables.
	// os.Getenv retrieves the value of the environment variable specified by the key.
	// In this case, it retrieves the value of the PORT environment variable.
	cfg := Config{
		Port: os.Getenv("PORT"),
	}

	// Return a pointer to the Config struct and nil error.
	// The caller can check the error value to determine if any errors occurred during configuration loading.
	return &cfg, nil
}
