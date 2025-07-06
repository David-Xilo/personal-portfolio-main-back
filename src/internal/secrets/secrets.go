package secrets

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
)

const JwtSecretName = "safehouse-jwt-signing-key"
const FrontendAuthSecretName = "safehouse-frontend-auth-key"

type SecretManager struct {
	client    *secretmanager.Client
	projectID string
}

type AppSecrets struct {
	JWTSigningKey   string
	FrontendAuthKey string
}

func NewSecretManager(ctx context.Context) (*SecretManager, error) {
	projectID := os.Getenv("GCP_PROJECT_ID")
	if projectID == "" {
		return nil, fmt.Errorf("GCP_PROJECT_ID environment variable not set")
	}

	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create secret manager client: %w", err)
	}

	return &SecretManager{
		client:    client,
		projectID: projectID,
	}, nil
}

func (sm *SecretManager) Close() error {
	return sm.client.Close()
}

func (sm *SecretManager) getSecret(ctx context.Context, secretName string) (string, error) {
	req := &secretmanagerpb.AccessSecretVersionRequest{
		Name: fmt.Sprintf("projects/%s/secrets/%s/versions/latest", sm.projectID, secretName),
	}

	result, err := sm.client.AccessSecretVersion(ctx, req)
	if err != nil {
		return "", fmt.Errorf("failed to access secret %s: %w", secretName, err)
	}

	return string(result.Payload.Data), nil
}

func (sm *SecretManager) LoadAppSecrets(ctx context.Context) (*AppSecrets, error) {
	slog.Info("Loading secrets from Google Cloud Secret Manager")

	jwtKey, err := sm.getSecret(ctx, JwtSecretName)
	if err != nil {
		return nil, fmt.Errorf("failed to load JWT signing key: %w", err)
	}

	frontendKey, err := sm.getSecret(ctx, FrontendAuthSecretName)
	if err != nil {
		return nil, fmt.Errorf("failed to load frontend auth key: %w", err)
	}

	slog.Info("Successfully loaded secrets from Google Cloud Secret Manager")

	return &AppSecrets{
		JWTSigningKey:   jwtKey,
		FrontendAuthKey: frontendKey,
	}, nil
}
