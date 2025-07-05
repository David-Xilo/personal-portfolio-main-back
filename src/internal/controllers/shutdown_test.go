package controllers

import (
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	configuration "safehouse-main-back/src/internal/config"
	"safehouse-main-back/src/internal/secrets"
	security2 "safehouse-main-back/src/internal/security"
)

func TestRouterSetup_GracefulShutdown(t *testing.T) {
	gin.SetMode(gin.TestMode)
	
	mockDB := new(MockDatabase)
	mockSecrets := &secrets.AppSecrets{
		JWTSigningKey:   "test-jwt-key",
		FrontendAuthKey: "test-auth-key",
	}
	config := configuration.Config{
		Environment:          "test",
		EnableHTTPSRedirect:  false,
		Port:                 "4000",
		FrontendURL:          "http://localhost:3000",
		DatabaseTimeout:      10 * time.Second,
		ReadTimeout:          10 * time.Second,
		WriteTimeout:         1 * time.Second,
		JWTSigningKey:        mockSecrets.JWTSigningKey,
		FrontendAuthKey:      mockSecrets.FrontendAuthKey,
		JWTExpirationMinutes: 30,
	}
	jwtManager := security2.NewJWTManager(config)
	
	// Create router setup
	routerSetup := SetupRoutes(mockDB, config, jwtManager)
	
	// Verify rate limiter is running
	assert.NotNil(t, routerSetup.RateLimiter)
	
	// Verify cleanup context is not cancelled initially
	select {
	case <-routerSetup.RateLimiter.GetCleanupContext().Done():
		t.Fatal("Cleanup context should not be cancelled initially")
	default:
		// Expected
	}
	
	// Test graceful shutdown
	routerSetup.RateLimiter.Stop()
	
	// Verify cleanup context is cancelled after Stop()
	select {
	case <-routerSetup.RateLimiter.GetCleanupContext().Done():
		// Expected
	case <-time.After(100 * time.Millisecond):
		t.Fatal("Cleanup context should be cancelled after Stop()")
	}
}