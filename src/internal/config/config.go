package configuration

import "os"

type Config struct {
	Environment         string
	EnableHTTPSRedirect bool
	Port                string
	FrontendURL         string
}

func LoadConfig() Config {
	env := getEnvOrDefault("ENV", "development")

	return Config{
		Environment:         env,
		EnableHTTPSRedirect: env == "production",
		FrontendURL:         getEnvOrDefault("FRONTEND_URL", "http://localhost:3000"),
	}
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
