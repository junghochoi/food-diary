package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"food-diary/internal/config"
	"food-diary/internal/handlers"
	"food-diary/internal/initializers"
	"food-diary/internal/repositories/pgsql"
	"food-diary/internal/routes"
	"food-diary/internal/services"
)

const version = "1.0.0"

func main() {
	var err error
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// Initialize Environment Variables and Database Connection
	initializers.LoadEnvVariables(logger)

	// Initialize Configuration
	conf, err := config.LoadConfig()

	// Initialize Database Connection
	db := initializers.ConnectToDB(logger)

	// DEPENDENCIES
	entryRepo := pgsql.NewEntryRepository(db)
	entryService := service.NewEntryService(entryRepo)
	h := handlers.NewHandlers(conf, db, entryService)
	r := routes.InitializeRoutes(h)

	// Start the Server
	var server *http.Server = &http.Server{
		Addr:         fmt.Sprintf(":%d", conf.Port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("Start %s server on %s", conf.Environment, server.Addr)
	err = server.ListenAndServe()
	logger.Fatal(err)
}
