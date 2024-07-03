package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Config(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("failed to loading .env file")
		os.Exit(1)
	}
	return os.Getenv(key)
}