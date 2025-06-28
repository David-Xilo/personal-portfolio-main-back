package configuration

import (
	"os"
	"time"
)

type Config struct {
	Environment         string
	EnableHTTPSRedirect bool
	Port                string
	FrontendURL         string
	SelfURL             string
	DatabaseTimeout     time.Duration
	ReadTimeout         time.Duration
	WriteTimeout        time.Duration
}

func LoadConfig() Config {
	env := getEnvOrDefault("ENV", "development")

	isProd := env == "production"

	frontendURL := getEnvOrDefault("FRONTEND_URL", "http://localhost:3000")
	port := getEnvOrDefault("SELF_URL", "4000")
	selfURL := getEnvOrDefault("SELF_URL", "http://localhost:"+port)

	dbTimeout, _ := time.ParseDuration(getEnvOrDefault("DATABASE_TIMEOUT", "10s"))
	readTimeout, _ := time.ParseDuration(getEnvOrDefault("READ_TIMEOUT", "10s"))
	// I don't have writes at the moment, used to init the server
	writeTimeout, _ := time.ParseDuration(getEnvOrDefault("WRITE_TIMEOUT", "1s"))

	return Config{
		Environment:         env,
		EnableHTTPSRedirect: isProd,
		FrontendURL:         frontendURL,
		SelfURL:             selfURL,
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
