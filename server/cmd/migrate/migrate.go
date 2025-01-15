package main

import (
	"log"
	"os"

	"food-diary/internal/initializers"
)

func main() {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	initializers.LoadEnvVariables(logger)
	db := initializers.ConnectToDB(logger)

	logger.Println(db)
}
