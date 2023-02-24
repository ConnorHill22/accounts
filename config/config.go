package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config func to get env value
func Config(key string) string {
	// load .env file
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("Error loading .env file, with error=%v", err)
	}
	// Return the value of the variable
	return os.Getenv(key)
}
