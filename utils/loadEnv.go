package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if _, err := os.Stat(".env"); err == nil {
		// Only load it if it exists
		if err := godotenv.Load(); err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}
		log.Println("Loaded environment variables from .env")
	} else {
		log.Println(".env file not found, assuming environment variables are set externally (Docker, etc.)")
	}
}
