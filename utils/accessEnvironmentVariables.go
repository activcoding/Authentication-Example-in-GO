package utils

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func GetVariable(key string) string {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
		return ""
	}
	signingKey := os.Getenv(key)
	return signingKey
}
