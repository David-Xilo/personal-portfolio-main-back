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
	DatabaseTimeout     time.Duration
}

func LoadConfig() Config {
	env := getEnvOrDefault("ENV", "development")

	dbTimeout, _ := time.ParseDuration(getEnvOrDefault("DATABASE_TIMEOUT", "10s"))

	return Config{
		Environment:         env,
		EnableHTTPSRedirect: env == "production",
		FrontendURL:         getEnvOrDefault("FRONTEND_URL", "http://localhost:3000"),
		DatabaseTimeout:     dbTimeout,
	}
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
