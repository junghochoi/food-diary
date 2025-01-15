package config

import (
	"errors"
	"flag"
	"os"
	"strconv"
)

type Config struct {
	Environment string
	AppName     string
	Port        int
	AppVersion  string
	JWTSecret   string
}

func LoadConfig() (*Config, error) {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		return nil, err
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return nil, errors.New("JWT_SECRET was not properly loaded")
	}

	conf := &Config{
		Port:        port,
		Environment: getEnv("ENV", "development"),
		AppVersion:  "1.0.0",
		JWTSecret:   jwtSecret,
	}

	flag.IntVar(&conf.Port, "port", port, "API server port")
	flag.StringVar(
		&conf.Environment,
		"env",
		"development",
		"Environment (development|staging|production)",
	)
	flag.Parse()

	return conf, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
