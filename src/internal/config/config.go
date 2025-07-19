package configuration

import (
	"fmt"
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
	DatabaseConfig       DbConfig
	ReadTimeout          time.Duration
	WriteTimeout         time.Duration
	JWTSigningKey        string
	FrontendAuthKey      string
	JWTExpirationMinutes int
}

type DbConfig struct {
	DbUrl      string
	UseIAMAuth bool
	DbHost     string
	DbName     string
	DbUser     string
	DbPort     string
	DbTimeout  time.Duration
}

func LoadConfig() Config {
	env := GetEnvOrDefault("ENV", "development")

	isProd := env == "production"

	useIAMAuthStr := GetEnvOrDefault("USE_IAM_DB_AUTH", "false")
	useIAMAuth := useIAMAuthStr == "true"

	frontendURL := GetEnvOrDefault("FRONTEND_URL", "http://localhost:80")
	port := GetEnvOrDefault("PORT", "8080")

	dbHost := GetEnvOrDefault("DB_HOST", "postgres-dev")
	dbUser := GetEnvOrDefault("DB_USER", "dev_user")
	dbName := GetEnvOrDefault("DB_NAME", "dev_db")
	dbPassword := GetEnvOrDefault("DB_PASSWORD", "dev_password")

	dbPortStr := GetEnvOrDefault("DB_PORT", "5432")

	defaultDbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser,
		dbPassword,
		dbHost,
		dbPortStr,
		dbName)
	dburl := GetEnvOrDefault("DATABASE_URL", defaultDbUrl)

	_, err := strconv.Atoi(dbPortStr)
	if err != nil {
		slog.Warn("Invalid DB_PORT value, falling back to default", "default", "5432")
	}

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

	jwtSigning := GetEnvOrDefault("JWT_SIGNING_KEY", "dev_jwt_signing_key")

	dbConfig := DbConfig{
		DbUrl:      dburl,
		UseIAMAuth: useIAMAuth,
		DbHost:     dbHost,
		DbName:     dbName,
		DbUser:     dbUser,
		DbPort:     dbPortStr,
		DbTimeout:  dbTimeout,
	}

	return Config{
		Environment:          env,
		EnableHTTPSRedirect:  isProd,
		FrontendURL:          frontendURL,
		Port:                 port,
		DatabaseConfig:       dbConfig,
		ReadTimeout:          readTimeout,
		WriteTimeout:         writeTimeout,
		JWTSigningKey:        jwtSigning,
		FrontendAuthKey:      secrets.FrontendTokenAuth,
		JWTExpirationMinutes: jwtExpiration,
	}
}

func GetEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func (config *Config) IsProduction() bool {
	return config.Environment == "production"
}
