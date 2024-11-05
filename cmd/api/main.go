package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"food-diary/internal/config"
	"food-diary/internal/initializers"
	"food-diary/internal/routes"
)

const version = "1.0.0"

func main() {
	var err error
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// Initialize Environment Variables and Database Connection
	initializers.LoadEnvVariables(logger)
	db := initializers.ConnectToDB(logger)

	// Initialize App Configs
	port, err := strconv.Atoi(os.Getenv("PORT"))
	var conf config.Config

	flag.IntVar(&conf.Port, "port", port, "API server port")
	flag.StringVar(
		&conf.Environment,
		"env",
		"development",
		"Environment (development|staging|production)",
	)
	conf.AppVersion = version
	flag.Parse()

	// Start the Server
	var server *http.Server = &http.Server{
		Addr:         fmt.Sprintf(":%d", conf.Port),
		Handler:      routes.InitializeRoutes(&conf, db),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("Start %s server on %s", conf.Environment, server.Addr)
	err = server.ListenAndServe()
	logger.Fatal(err)
}
