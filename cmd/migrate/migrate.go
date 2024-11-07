package main

import (
	"log"
	"os"

	"food-diary/internal/initializers"
	"food-diary/internal/models"
)

func main() {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	initializers.LoadEnvVariables(logger)
	db := initializers.ConnectToDB(logger)

	db.AutoMigrate(&models.Entry{})

	// db.Migrator().AlterColumn(&models.Entry{}, "Foods")
}
