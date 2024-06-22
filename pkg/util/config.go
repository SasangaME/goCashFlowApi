package util

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
)

func Config(key string) (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", errors.New("failed to load .env file")
	}
	return os.Getenv(key), nil
}
