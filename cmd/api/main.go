package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"food-diary/internal/config"
	"food-diary/internal/routes"
)

const version = "1.0.0"

func main() {
	var conf config.Config

	flag.IntVar(&conf.Port, "port", 4000, "API server port")
	flag.StringVar(
		&conf.Environment,
		"env",
		"development",
		"Environment (development|staging|production)",
	)
  conf.AppVersion = version 
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	var server *http.Server = &http.Server{
		Addr:         fmt.Sprintf(":%d", conf.Port),
		Handler:      routes.InitializeRoutes(&conf),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("starting %s server on %s", conf.Environment, server.Addr)
	err := server.ListenAndServe()
	logger.Fatal(err)
}
