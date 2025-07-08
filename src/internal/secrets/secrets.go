package secrets

import (
	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	"context"
	"fmt"
	"log/slog"
	"os"
)

// It doesnt make sense to have JWT without credentials, but I'm leaving it here in case I want to create credentials later
const JwtSecretName = "safehouse-jwt-signing-key"
const DbSecretName = "safehouse-db-password"
const FrontendTokenAuth = "safehouse-frontend"

type AppSecrets struct {
	JWTSigningKey string
	DbPassword    string
}

type SecretProvider interface {
	LoadAppSecrets(ctx context.Context) (*AppSecrets, error)
	Close() error
}

func NewSecretProvider(ctx context.Context) (SecretProvider, error) {
	secretMode := os.Getenv("ENV")

	switch secretMode {
	case "development":
		slog.Info("Using local secret provider for development")
		return NewLocalSecretProvider()
	case "production":
		slog.Info("Using GCP Secret Manager")
		return NewGCPSecretManager(ctx)
	default:
		return nil, fmt.Errorf("unknown ENV: %s", secretMode)
	}
}

type GCPSecretManager struct {
	client    *secretmanager.Client
	projectID string
}

func NewGCPSecretManager(ctx context.Context) (*GCPSecretManager, error) {
	projectID := os.Getenv("GCP_PROJECT_ID")
	if projectID == "" {
		return nil, fmt.Errorf("GCP_PROJECT_ID environment variable not set")
	}

	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create secret manager client: %w", err)
	}

	return &GCPSecretManager{
		client:    client,
		projectID: projectID,
	}, nil
}

func (sm *GCPSecretManager) Close() error {
	return sm.client.Close()
}

func (sm *GCPSecretManager) getSecret(ctx context.Context, secretName string) (string, error) {
	req := &secretmanagerpb.AccessSecretVersionRequest{
		Name: fmt.Sprintf("projects/%s/secrets/%s/versions/latest", sm.projectID, secretName),
	}

	result, err := sm.client.AccessSecretVersion(ctx, req)
	if err != nil {
		return "", fmt.Errorf("failed to access secret %s: %w", secretName, err)
	}

	return string(result.Payload.Data), nil
}

func (sm *GCPSecretManager) LoadAppSecrets(ctx context.Context) (*AppSecrets, error) {
	slog.Info("Loading secrets from Google Cloud Secret Manager")

	jwtKey, err := sm.getSecret(ctx, JwtSecretName)
	if err != nil {
		return nil, fmt.Errorf("failed to load JWT signing key: %w", err)
	}

	dbPassword, err := sm.getSecret(ctx, DbSecretName)
	if err != nil {
		return nil, fmt.Errorf("failed to load JWT signing key: %w", err)
	}

	slog.Info("Successfully loaded secrets from Google Cloud Secret Manager")

	return &AppSecrets{
		JWTSigningKey: jwtKey,
		DbPassword:    dbPassword,
	}, nil
}
