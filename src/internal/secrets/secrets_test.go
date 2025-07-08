package secrets

import (
	"context"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewSecretProvider(t *testing.T) {
	tests := []struct {
		name        string
		env         string
		expectError bool
		expectType  string
	}{
		{
			name:        "development environment returns LocalSecretProvider",
			env:         "development",
			expectError: false,
			expectType:  "*secrets.LocalSecretProvider",
		},
		{
			name:        "production environment returns GCPSecretManager",
			env:         "production",
			expectError: false,
			expectType:  "*secrets.GCPSecretManager",
		},
		{
			name:        "unknown environment returns error",
			env:         "unknown",
			expectError: true,
			expectType:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv("ENV", tt.env)
			defer os.Unsetenv("ENV")

			// For production environment, we need to set GCP_PROJECT_ID
			if tt.env == "production" {
				os.Setenv("GCP_PROJECT_ID", "test-project")
				defer os.Unsetenv("GCP_PROJECT_ID")
			}

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			provider, err := NewSecretProvider(ctx)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, provider)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, provider)
				if provider != nil {
					provider.Close()
				}
			}
		})
	}
}

func TestLocalSecretProvider_getSecret(t *testing.T) {
	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "secrets-test")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	// Create a test secret file
	secretName := "test-secret"
	secretValue := "secret-value"
	secretPath := filepath.Join(tempDir, secretName)
	err = os.WriteFile(secretPath, []byte(secretValue), 0644)
	require.NoError(t, err)

	tests := []struct {
		name        string
		secretsPath string
		secretName  string
		envValue    string
		expectError bool
		expectedVal string
	}{
		{
			name:        "reads secret from file",
			secretsPath: tempDir,
			secretName:  secretName,
			envValue:    "",
			expectError: false,
			expectedVal: secretValue,
		},
		{
			name:        "falls back to environment variable",
			secretsPath: "",
			secretName:  "test-env-secret",
			envValue:    "env-secret-value",
			expectError: false,
			expectedVal: "env-secret-value",
		},
		{
			name:        "returns error when secret not found",
			secretsPath: "",
			secretName:  "nonexistent-secret",
			envValue:    "",
			expectError: true,
			expectedVal: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set up environment variable if provided
			if tt.envValue != "" {
				envKey := "TEST_ENV_SECRET"
				os.Setenv(envKey, tt.envValue)
				defer os.Unsetenv(envKey)
			}

			provider := &LocalSecretProvider{
				secretsPath: tt.secretsPath,
			}

			result, err := provider.getSecret(tt.secretName)

			if tt.expectError {
				assert.Error(t, err)
				assert.Empty(t, result)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedVal, result)
			}
		})
	}
}

func TestLocalSecretProvider_LoadAppSecrets(t *testing.T) {
	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "secrets-test")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	// Create secret files
	jwtSecret := "test-jwt-key"
	dbPassword := "test-db-password"

	err = os.WriteFile(filepath.Join(tempDir, JwtSecretName), []byte(jwtSecret), 0644)
	require.NoError(t, err)

	err = os.WriteFile(filepath.Join(tempDir, DbSecretName), []byte(dbPassword), 0644)
	require.NoError(t, err)

	provider := &LocalSecretProvider{
		secretsPath: tempDir,
	}

	ctx := context.Background()
	secrets, err := provider.LoadAppSecrets(ctx)

	assert.NoError(t, err)
	assert.NotNil(t, secrets)
	assert.Equal(t, jwtSecret, secrets.JWTSigningKey)
	assert.Equal(t, dbPassword, secrets.DbPassword)
}

func TestLocalSecretProvider_LoadAppSecrets_EnvironmentFallback(t *testing.T) {
	// Set environment variables
	jwtSecret := "env-jwt-key"
	dbPassword := "env-db-password"

	os.Setenv("SAFEHOUSE_JWT_SIGNING_KEY", jwtSecret)
	os.Setenv("SAFEHOUSE_DB_PASSWORD", dbPassword)
	defer os.Unsetenv("SAFEHOUSE_JWT_SIGNING_KEY")
	defer os.Unsetenv("SAFEHOUSE_DB_PASSWORD")

	provider := &LocalSecretProvider{
		secretsPath: "", // No secrets directory
	}

	ctx := context.Background()
	secrets, err := provider.LoadAppSecrets(ctx)

	assert.NoError(t, err)
	assert.NotNil(t, secrets)
	assert.Equal(t, jwtSecret, secrets.JWTSigningKey)
	assert.Equal(t, dbPassword, secrets.DbPassword)
}

func TestLocalSecretProvider_LoadAppSecrets_MissingSecret(t *testing.T) {
	provider := &LocalSecretProvider{
		secretsPath: "",
	}

	ctx := context.Background()
	secrets, err := provider.LoadAppSecrets(ctx)

	assert.Error(t, err)
	assert.Nil(t, secrets)
	assert.Contains(t, err.Error(), "failed to load JWT signing key")
}

func TestNewLocalSecretProvider(t *testing.T) {
	tests := []struct {
		name        string
		secretsPath string
		expectPath  string
	}{
		{
			name:        "uses environment variable when set",
			secretsPath: "/custom/path",
			expectPath:  "/custom/path",
		},
		{
			name:        "uses default path when env not set",
			secretsPath: "",
			expectPath:  "/app/secrets",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Unsetenv("SECRETS_PATH")

			if tt.secretsPath != "" {
				os.Setenv("SECRETS_PATH", tt.secretsPath)
				defer os.Unsetenv("SECRETS_PATH")
			}

			provider, err := NewLocalSecretProvider()

			assert.NoError(t, err)
			assert.NotNil(t, provider)

			// Since the directory likely doesn't exist, it should be empty
			if _, err := os.Stat(tt.expectPath); os.IsNotExist(err) {
				assert.Empty(t, provider.secretsPath)
			} else {
				assert.Equal(t, tt.expectPath, provider.secretsPath)
			}
		})
	}
}

func TestSecretConstants(t *testing.T) {
	assert.Equal(t, "safehouse-jwt-signing-key", JwtSecretName)
	assert.Equal(t, "safehouse-db-password", DbSecretName)
	assert.Equal(t, "safehouse-frontend", FrontendTokenAuth)
}

func TestAppSecrets_Struct(t *testing.T) {
	secrets := &AppSecrets{
		JWTSigningKey: "test-jwt",
		DbPassword:    "test-password",
	}

	assert.Equal(t, "test-jwt", secrets.JWTSigningKey)
	assert.Equal(t, "test-password", secrets.DbPassword)
}
