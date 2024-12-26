package initializers

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables(logger *log.Logger, environment string) {

	envFile := fmt.Sprintf(".%s.env", environment)
	logger.Println("Initialize Environment Variables")
	logger.Printf("Loading environment variables from %s", envFile)

	// Determine the .env file based on the environment

	// Load the environment-specific file
	err := godotenv.Load(envFile)
	if err != nil {
		logger.Fatalf("Failed to load environment variables from %s: %v", envFile, err)
	}

}
