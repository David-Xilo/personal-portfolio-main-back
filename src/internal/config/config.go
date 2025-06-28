package configuration

import (
	"log"
	"os"
	"time"
)

type Config struct {
	Environment         string
	EnableHTTPSRedirect bool
	Port                string
	FrontendURL         string
	DatabaseTimeout     time.Duration
	ReadTimeout         time.Duration
	WriteTimeout        time.Duration
}

func LoadConfig() Config {
	env := getEnvOrDefault("ENV", "development")

	isProd := env == "production"

	frontendURL := getEnvOrDefault("FRONTEND_URL", "http://localhost:3000")
	port := getEnvOrDefault("PORT", "4000")

	dbTimeoutStr := getEnvOrDefault("DATABASE_TIMEOUT", "10s")
	dbTimeout, err := time.ParseDuration(getEnvOrDefault("DATABASE_TIMEOUT", "10s"))
	if err != nil {
		log.Printf("Invalid DATABASE_TIMEOUT value '%s', falling back to default: 10s.", dbTimeoutStr)
		dbTimeout = 10 * time.Second
	}

	readTimeoutStr := getEnvOrDefault("READ_TIMEOUT", "10s")
	readTimeout, err := time.ParseDuration(readTimeoutStr)
	if err != nil {
		log.Printf("Invalid READ_TIMEOUT value '%s', falling back to default: 10s.", readTimeoutStr)
		readTimeout = 10 * time.Second
	}
	// I don't have writes at the moment, used to init the server
	writeTimeoutStr := getEnvOrDefault("WRITE_TIMEOUT", "1s")
	writeTimeout, err := time.ParseDuration(writeTimeoutStr)
	if err != nil {
		log.Printf("Invalid WRITE_TIMEOUT value '%s', falling back to default: 1s.", writeTimeoutStr)
		writeTimeout = 1 * time.Second
	}

	return Config{
		Environment:         env,
		EnableHTTPSRedirect: isProd,
		FrontendURL:         frontendURL,
		Port:                port,
		DatabaseTimeout:     dbTimeout,
		ReadTimeout:         readTimeout,
		WriteTimeout:        writeTimeout,
	}
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
