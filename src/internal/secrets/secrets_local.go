package secrets

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
)

type LocalSecretProvider struct {
	secretsPath string
}

func NewLocalSecretProvider() (*LocalSecretProvider, error) {
	secretsPath := os.Getenv("SECRETS_PATH")
	if secretsPath == "" {
		secretsPath = "/app/secrets" // Default path when mounted in container
	}

	// Check if secrets directory exists, if not, use environment variables only
	if _, err := os.Stat(secretsPath); os.IsNotExist(err) {
		slog.Info("Secrets directory not found, using environment variables only")
		secretsPath = ""
	} else {
		slog.Info("Using secrets directory")
	}

	return &LocalSecretProvider{
		secretsPath: secretsPath,
	}, nil
}

func (lsp *LocalSecretProvider) Close() error {
	// No cleanup needed for local provider
	return nil
}

func (lsp *LocalSecretProvider) getSecret(secretName string) (string, error) {
	// Try to read from file first
	if lsp.secretsPath != "" {
		filePath := filepath.Join(lsp.secretsPath, secretName)
		if content, err := os.ReadFile(filePath); err == nil {
			secret := strings.TrimSpace(string(content))
			if secret != "" {
				slog.Info("Loaded secret from file")
				return secret, nil
			}
		} else {
			slog.Info("Could not read secret from file")
		}
	}

	// Fallback to environment variable
	envKey := strings.ToUpper(strings.ReplaceAll(secretName, "-", "_"))
	if value := os.Getenv(envKey); value != "" {
		slog.Info("Loaded secret from environment variable")
		return value, nil
	}

	return "", fmt.Errorf("secret not found in files or environment variables")
}

func (lsp *LocalSecretProvider) LoadAppSecrets(ctx context.Context) (*AppSecrets, error) {
	slog.Info("Loading secrets from local provider (files/environment)")

	jwtKey, err := lsp.getSecret(JwtSecretName)
	if err != nil {
		return nil, fmt.Errorf("failed to load JWT signing key: %w", err)
	}

	frontendKey, err := lsp.getSecret(FrontendAuthSecretName)
	if err != nil {
		return nil, fmt.Errorf("failed to load frontend auth key: %w", err)
	}

	slog.Info("Successfully loaded secrets from local provider")

	return &AppSecrets{
		JWTSigningKey:   jwtKey,
		FrontendAuthKey: frontendKey,
	}, nil
}
