package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found or failed to load.")
	}
}

func GetEnv(key string) string {
	return os.Getenv(key)
}
