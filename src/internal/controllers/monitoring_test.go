package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/time/rate"
	configuration "safehouse-main-back/src/internal/config"
	"safehouse-main-back/src/internal/middleware"
	"safehouse-main-back/src/internal/secrets"
	security2 "safehouse-main-back/src/internal/security"
)

func TestMonitoringEndpoints_RateLimiterStats(t *testing.T) {
	gin.SetMode(gin.TestMode)
	
	// Create a rate limiter
	rateLimiter := middleware.NewIPRateLimiter(rate.Limit(10), 5)
	defer rateLimiter.Stop()
	
	// Create router with monitoring endpoints
	router := gin.New()
	addMonitoringEndpoints(router, rateLimiter)
	
	// Add some test data to rate limiter
	rateLimiter.GetLimiter("192.168.1.1")
	rateLimiter.GetLimiter("192.168.1.2")
	
	// Test rate limiter stats endpoint
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/internal/stats/rate-limiter", nil)
	router.ServeHTTP(w, req)
	
	assert.Equal(t, http.StatusOK, w.Code)
	
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)
	
	// Verify response structure
	assert.Contains(t, response, "rate_limiter")
	assert.Contains(t, response, "timestamp")
	
	rateLimiterStats := response["rate_limiter"].(map[string]interface{})
	assert.Contains(t, rateLimiterStats, "total_ips")
	assert.Contains(t, rateLimiterStats, "limit")
	assert.Contains(t, rateLimiterStats, "burst")
	
	// Verify actual values
	assert.Equal(t, float64(2), rateLimiterStats["total_ips"])
	assert.Equal(t, float64(10), rateLimiterStats["limit"])
	assert.Equal(t, float64(5), rateLimiterStats["burst"])
	
	// Verify timestamp is recent
	timestamp := response["timestamp"].(float64)
	now := time.Now().Unix()
	assert.True(t, timestamp <= float64(now))
	assert.True(t, timestamp >= float64(now-5)) // Within 5 seconds
}

func TestSetupRoutes_IncludesMonitoringEndpoints(t *testing.T) {
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
	
	routerSetup := SetupRoutes(mockDB, config, jwtManager)
	defer routerSetup.RateLimiter.Stop()
	
	// Test that monitoring endpoint is available
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/internal/stats/rate-limiter", nil)
	routerSetup.Router.ServeHTTP(w, req)
	
	assert.Equal(t, http.StatusOK, w.Code)
	
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)
	
	assert.Contains(t, response, "rate_limiter")
}