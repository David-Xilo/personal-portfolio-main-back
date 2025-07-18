package configuration

import (
	"os"
	"safehouse-main-back/src/internal/secrets"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetEnvOrDefault(t *testing.T) {
	tests := []struct {
		name         string
		key          string
		defaultValue string
		envValue     string
		expected     string
	}{
		{
			name:         "returns environment value when set",
			key:          "TEST_KEY",
			defaultValue: "default",
			envValue:     "env_value",
			expected:     "env_value",
		},
		{
			name:         "returns default when environment not set",
			key:          "TEST_KEY_NOT_SET",
			defaultValue: "default",
			envValue:     "",
			expected:     "default",
		},
		{
			name:         "returns default when environment is empty",
			key:          "TEST_KEY_EMPTY",
			defaultValue: "default",
			envValue:     "",
			expected:     "default",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Clean up before test
			os.Unsetenv(tt.key)

			// Set environment variable if provided
			if tt.envValue != "" {
				os.Setenv(tt.key, tt.envValue)
				defer os.Unsetenv(tt.key)
			}

			result := GetEnvOrDefault(tt.key, tt.defaultValue)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestLoadConfig_DefaultValues(t *testing.T) {
	// Clear all relevant environment variables
	envVars := []string{"ENV", "FRONTEND_URL", "PORT", "DATABASE_TIMEOUT", "READ_TIMEOUT", "WRITE_TIMEOUT"}
	for _, env := range envVars {
		os.Unsetenv(env)
	}

	mockSecrets := &secrets.AppSecrets{
		JWTSigningKey: "test-jwt-key",
		DbPassword:    "test-db-password",
	}
	config := LoadConfig(mockSecrets)

	assert.Equal(t, "development", config.Environment)
	assert.False(t, config.EnableHTTPSRedirect)
	assert.Equal(t, "http://localhost:80", config.FrontendURL)
	assert.Equal(t, "8080", config.Port)
	assert.Equal(t, 10*time.Second, config.DatabaseConfig.DbTimeout)
	assert.Equal(t, "postgres-dev", config.DatabaseConfig.DbHost)
	assert.Equal(t, "dev_user", config.DatabaseConfig.DbUser)
	assert.Equal(t, "dev_db", config.DatabaseConfig.DbName)
	assert.Equal(t, "5432", config.DatabaseConfig.DbPort)
	// DbPassword is no longer part of DbConfig struct - it's used internally in LoadConfig
	assert.Equal(t, 10*time.Second, config.ReadTimeout)
	assert.Equal(t, 1*time.Second, config.WriteTimeout)
}

func TestLoadConfig_ProductionEnvironment(t *testing.T) {
	// Set production environment
	os.Setenv("ENV", "production")
	defer os.Unsetenv("ENV")

	mockSecrets := &secrets.AppSecrets{
		JWTSigningKey: "test-jwt-key",
		DbPassword:    "test-db-password",
	}
	config := LoadConfig(mockSecrets)

	assert.Equal(t, "production", config.Environment)
	assert.True(t, config.EnableHTTPSRedirect)
}

func TestLoadConfig_CustomValues(t *testing.T) {
	// Set custom environment variables
	envVars := map[string]string{
		"ENV":              "staging",
		"FRONTEND_URL":     "https://example.com",
		"PORT":             "8080",
		"DATABASE_TIMEOUT": "30s",
		"READ_TIMEOUT":     "15s",
		"WRITE_TIMEOUT":    "5s",
	}

	for key, value := range envVars {
		os.Setenv(key, value)
		defer os.Unsetenv(key)
	}

	mockSecrets := &secrets.AppSecrets{
		JWTSigningKey: "test-jwt-key",
		DbPassword:    "test-db-password",
	}
	config := LoadConfig(mockSecrets)

	assert.Equal(t, "staging", config.Environment)
	assert.False(t, config.EnableHTTPSRedirect) // Only "production" enables HTTPS redirect
	assert.Equal(t, "https://example.com", config.FrontendURL)
	assert.Equal(t, "8080", config.Port)
	assert.Equal(t, 30*time.Second, config.DatabaseConfig.DbTimeout)
	assert.Equal(t, 15*time.Second, config.ReadTimeout)
	assert.Equal(t, 5*time.Second, config.WriteTimeout)
}

func TestLoadConfig_InvalidTimeouts(t *testing.T) {
	tests := []struct {
		name     string
		envVar   string
		value    string
		expected time.Duration
	}{
		{
			name:     "invalid database timeout falls back to default",
			envVar:   "DATABASE_TIMEOUT",
			value:    "invalid",
			expected: 10 * time.Second,
		},
		{
			name:     "invalid read timeout falls back to default",
			envVar:   "READ_TIMEOUT",
			value:    "not-a-duration",
			expected: 10 * time.Second,
		},
		{
			name:     "invalid write timeout falls back to default",
			envVar:   "WRITE_TIMEOUT",
			value:    "xyz",
			expected: 1 * time.Second,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Clear all timeout environment variables first
			timeoutVars := []string{"DATABASE_TIMEOUT", "READ_TIMEOUT", "WRITE_TIMEOUT"}
			for _, env := range timeoutVars {
				os.Unsetenv(env)
			}

			// Set the specific invalid value
			os.Setenv(tt.envVar, tt.value)
			defer os.Unsetenv(tt.envVar)

			mockSecrets := &secrets.AppSecrets{
				JWTSigningKey: "test-jwt-key",
				DbPassword:    "test-db-password",
			}
			config := LoadConfig(mockSecrets)

			switch tt.envVar {
			case "DATABASE_TIMEOUT":
				assert.Equal(t, tt.expected, config.DatabaseConfig.DbTimeout)
			case "READ_TIMEOUT":
				assert.Equal(t, tt.expected, config.ReadTimeout)
			case "WRITE_TIMEOUT":
				assert.Equal(t, tt.expected, config.WriteTimeout)
			}
		})
	}
}

func TestLoadConfig_ValidTimeouts(t *testing.T) {
	tests := []struct {
		name     string
		envVar   string
		value    string
		expected time.Duration
	}{
		{
			name:     "valid database timeout in seconds",
			envVar:   "DATABASE_TIMEOUT",
			value:    "25s",
			expected: 25 * time.Second,
		},
		{
			name:     "valid read timeout in minutes",
			envVar:   "READ_TIMEOUT",
			value:    "2m",
			expected: 2 * time.Minute,
		},
		{
			name:     "valid write timeout in milliseconds",
			envVar:   "WRITE_TIMEOUT",
			value:    "500ms",
			expected: 500 * time.Millisecond,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Clear all timeout environment variables first
			timeoutVars := []string{"DATABASE_TIMEOUT", "READ_TIMEOUT", "WRITE_TIMEOUT"}
			for _, env := range timeoutVars {
				os.Unsetenv(env)
			}

			// Set the specific valid value
			os.Setenv(tt.envVar, tt.value)
			defer os.Unsetenv(tt.envVar)

			mockSecrets := &secrets.AppSecrets{
				JWTSigningKey: "test-jwt-key",
				DbPassword:    "test-db-password",
			}
			config := LoadConfig(mockSecrets)

			switch tt.envVar {
			case "DATABASE_TIMEOUT":
				assert.Equal(t, tt.expected, config.DatabaseConfig.DbTimeout)
			case "READ_TIMEOUT":
				assert.Equal(t, tt.expected, config.ReadTimeout)
			case "WRITE_TIMEOUT":
				assert.Equal(t, tt.expected, config.WriteTimeout)
			}
		})
	}
}

func TestConfig_Struct(t *testing.T) {
	config := Config{
		Environment:         "test",
		EnableHTTPSRedirect: true,
		Port:                "80",
		FrontendURL:         "http://test.com",
		DatabaseConfig: DbConfig{
			DbHost:     "test-host",
			DbName:     "test-db",
			DbUser:     "test-user",
			DbPort:     "5432",
			// DbPassword removed from DbConfig struct
			DbTimeout:  5 * time.Second,
		},
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 2 * time.Second,
	}

	assert.Equal(t, "test", config.Environment)
	assert.True(t, config.EnableHTTPSRedirect)
	assert.Equal(t, "80", config.Port)
	assert.Equal(t, "http://test.com", config.FrontendURL)
	assert.Equal(t, 5*time.Second, config.DatabaseConfig.DbTimeout)
	assert.Equal(t, 15*time.Second, config.ReadTimeout)
	assert.Equal(t, 2*time.Second, config.WriteTimeout)
}

func TestLoadConfig_DatabaseConfig(t *testing.T) {
	// Set custom database environment variables
	envVars := map[string]string{
		"DB_HOST":          "custom-host",
		"DB_USER":          "custom-user",
		"DB_NAME":          "custom-db",
		"DB_PORT":          "3306",
		"DATABASE_TIMEOUT": "20s",
	}

	for key, value := range envVars {
		os.Setenv(key, value)
		defer os.Unsetenv(key)
	}

	mockSecrets := &secrets.AppSecrets{
		JWTSigningKey: "test-jwt-key",
		DbPassword:    "secret-password",
	}
	config := LoadConfig(mockSecrets)

	assert.Equal(t, "custom-host", config.DatabaseConfig.DbHost)
	assert.Equal(t, "custom-user", config.DatabaseConfig.DbUser)
	assert.Equal(t, "custom-db", config.DatabaseConfig.DbName)
	assert.Equal(t, "3306", config.DatabaseConfig.DbPort)
	// DbPassword is no longer part of DbConfig struct - it's used internally in LoadConfig
	assert.Equal(t, 20*time.Second, config.DatabaseConfig.DbTimeout)
}

func TestLoadConfig_InvalidDBPort(t *testing.T) {
	// Set invalid DB port
	os.Setenv("DB_PORT", "invalid-port")
	defer os.Unsetenv("DB_PORT")

	mockSecrets := &secrets.AppSecrets{
		JWTSigningKey: "test-jwt-key",
		DbPassword:    "test-password",
	}
	config := LoadConfig(mockSecrets)

	// Should still use the invalid port string but log warning
	assert.Equal(t, "invalid-port", config.DatabaseConfig.DbPort)
}

func TestLoadConfig_JWTExpiration(t *testing.T) {
	tests := []struct {
		name     string
		envValue string
		expected int
	}{
		{
			name:     "valid JWT expiration",
			envValue: "60",
			expected: 60,
		},
		{
			name:     "invalid JWT expiration falls back to default",
			envValue: "invalid",
			expected: 30,
		},
		{
			name:     "empty JWT expiration uses default",
			envValue: "",
			expected: 30,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Unsetenv("JWT_EXPIRATION_MINUTES")

			if tt.envValue != "" {
				os.Setenv("JWT_EXPIRATION_MINUTES", tt.envValue)
				defer os.Unsetenv("JWT_EXPIRATION_MINUTES")
			}

			mockSecrets := &secrets.AppSecrets{
				JWTSigningKey: "test-jwt-key",
				DbPassword:    "test-password",
			}
			config := LoadConfig(mockSecrets)

			assert.Equal(t, tt.expected, config.JWTExpirationMinutes)
		})
	}
}
