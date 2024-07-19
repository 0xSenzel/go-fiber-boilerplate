package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// init function runs automatically when the package is imported
func init() {
	err := godotenv.Load("./configs/.env")
	if err != nil {
		panic("Error loading .env file: " + err.Error())
	}
}

// LoadEnv retrieves the value of the environment variable named by the key.
// It returns the value and an error if the environment variable is not set.
func LoadEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Environment variable not set: %s", key)
	}
	return value
}