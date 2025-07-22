package controllers

import (
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	configuration "personal-portfolio-main-back/src/internal/config"
	security2 "personal-portfolio-main-back/src/internal/security"
)

func TestRouterSetup_GracefulShutdown(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockDB := new(MockDatabase)
	config := configuration.Config{
		Environment:         "test",
		EnableHTTPSRedirect: false,
		Port:                "4000",
		AllowedOrigins:      []string{"http://localhost:3000"},
		DatabaseConfig: configuration.DbConfig{
			DbTimeout: 10 * time.Second,
		},
		ReadTimeout:          10 * time.Second,
		WriteTimeout:         1 * time.Second,
		JWTSigningKey:        "JWTSigningKey",
		FrontendAuthKey:      configuration.FrontendTokenAuth,
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
