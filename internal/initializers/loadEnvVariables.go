package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables(logger *log.Logger) {
	logger.Println("Initialize Environment Varaibles")
	err := godotenv.Load()
	if err != nil {
		logger.Fatal("Failed to Load Environment Varibales")
	}
}
