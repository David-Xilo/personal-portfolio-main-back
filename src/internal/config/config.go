package configuration

import (
	"log/slog"
	"os"
	"safehouse-main-back/src/internal/secrets"
	"strconv"
	"time"
)

type Config struct {
	Environment          string
	EnableHTTPSRedirect  bool
	Port                 string
	FrontendURL          string
	DatabaseTimeout      time.Duration
	ReadTimeout          time.Duration
	WriteTimeout         time.Duration
	JWTSigningKey        string
	FrontendAuthKey      string
	JWTExpirationMinutes int
}

func LoadConfig(appSecrets *secrets.AppSecrets) Config {
	env := GetEnvOrDefault("ENV", "development")

	isProd := env == "production"

	frontendURL := GetEnvOrDefault("FRONTEND_URL", "http://localhost:3000")
	port := GetEnvOrDefault("PORT", "4000")

	dbTimeoutStr := GetEnvOrDefault("DATABASE_TIMEOUT", "10s")
	dbTimeout, err := time.ParseDuration(dbTimeoutStr)
	if err != nil {
		slog.Warn("Invalid DATABASE_TIMEOUT value, falling back to default", "default", "10s")
		dbTimeout = 10 * time.Second
	}

	readTimeoutStr := GetEnvOrDefault("READ_TIMEOUT", "10s")
	readTimeout, err := time.ParseDuration(readTimeoutStr)
	if err != nil {
		slog.Warn("Invalid READ_TIMEOUT value, falling back to default", "default", "10s")
		readTimeout = 10 * time.Second
	}

	// I don't have writes at the moment, used to init the server
	writeTimeoutStr := GetEnvOrDefault("WRITE_TIMEOUT", "1s")
	writeTimeout, err := time.ParseDuration(writeTimeoutStr)
	if err != nil {
		slog.Warn("Invalid WRITE_TIMEOUT value, falling back to default", "default", "1s")
		writeTimeout = 1 * time.Second
	}

	jwtExpirationStr := GetEnvOrDefault("JWT_EXPIRATION_MINUTES", "30")
	jwtExpiration, err := strconv.Atoi(jwtExpirationStr)
	if err != nil {
		slog.Warn("Invalid JWT_EXPIRATION_MINUTES value, falling back to default", "default", "30")
		jwtExpiration = 30
	}

	return Config{
		Environment:          env,
		EnableHTTPSRedirect:  isProd,
		FrontendURL:          frontendURL,
		Port:                 port,
		DatabaseTimeout:      dbTimeout,
		ReadTimeout:          readTimeout,
		WriteTimeout:         writeTimeout,
		JWTSigningKey:        appSecrets.JWTSigningKey,
		FrontendAuthKey:      appSecrets.FrontendAuthKey,
		JWTExpirationMinutes: jwtExpiration,
	}
}

func GetEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
