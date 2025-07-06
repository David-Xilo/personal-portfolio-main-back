package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"safehouse-main-back/src/internal/secrets"
	security2 "safehouse-main-back/src/internal/security"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	configuration "safehouse-main-back/src/internal/config"
	"safehouse-main-back/src/internal/models"
)

// MockDatabase implements the Database interface for testing
type MockDatabase struct {
	mock.Mock
}

func (m *MockDatabase) GetContact() (*models.Contacts, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Contacts), args.Error(1)
}

func (m *MockDatabase) GetProjects(projectType models.ProjectType) ([]*models.ProjectGroups, error) {
	args := m.Called(projectType)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*models.ProjectGroups), args.Error(1)
}

func (m *MockDatabase) GetGamesPlayed() ([]*models.GamesPlayed, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*models.GamesPlayed), args.Error(1)
}

func TestSetupRoutes(t *testing.T) {
	// Set test mode for gin
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

	assert.NotNil(t, routerSetup)
	assert.NotNil(t, routerSetup.Router)
	assert.NotNil(t, routerSetup.RateLimiter)

	// Test that the router is created and has expected behavior
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)
	routerSetup.Router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, "healthy", response["status"])
	assert.Contains(t, response, "timestamp")
}

func TestCreateRouter(t *testing.T) {
	gin.SetMode(gin.TestMode)

	config := configuration.Config{
		Environment:         "test",
		EnableHTTPSRedirect: false,
		Port:                "4000",
		FrontendURL:         "http://localhost:3000",
		DatabaseTimeout:     10 * time.Second,
		ReadTimeout:         10 * time.Second,
		WriteTimeout:        1 * time.Second,
	}

	routerSetup := createRouter(config)

	assert.NotNil(t, routerSetup)
	assert.NotNil(t, routerSetup.Router)
	assert.NotNil(t, routerSetup.RateLimiter)

	// Test basic functionality
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/nonexistent", nil)
	routerSetup.Router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusForbidden, w.Code)
}

func TestCreateRouter_WithHTTPSRedirect(t *testing.T) {
	gin.SetMode(gin.TestMode)

	config := configuration.Config{
		Environment:         "production",
		EnableHTTPSRedirect: true,
		Port:                "4000",
		FrontendURL:         "https://example.com",
		DatabaseTimeout:     10 * time.Second,
		ReadTimeout:         10 * time.Second,
		WriteTimeout:        1 * time.Second,
	}

	routerSetup := createRouter(config)

	assert.NotNil(t, routerSetup)
	assert.NotNil(t, routerSetup.Router)
	assert.NotNil(t, routerSetup.RateLimiter)
}

func TestGetControllers(t *testing.T) {
	mockDB := new(MockDatabase)
	config := configuration.Config{
		Environment:         "test",
		EnableHTTPSRedirect: false,
		Port:                "4000",
		FrontendURL:         "http://localhost:3000",
		DatabaseTimeout:     10 * time.Second,
		ReadTimeout:         10 * time.Second,
		WriteTimeout:        1 * time.Second,
	}

	mockSecrets := &secrets.AppSecrets{
		JWTSigningKey:   "test-jwt-key",
		FrontendAuthKey: "test-auth-key",
	}
	config.JWTSigningKey = mockSecrets.JWTSigningKey
	config.FrontendAuthKey = mockSecrets.FrontendAuthKey
	config.JWTExpirationMinutes = 30
	jwtManager := security2.NewJWTManager(config)

	controllers := getControllers(mockDB, config, jwtManager)

	assert.Len(t, controllers, 4) // about, tech, games, finance

	// Verify that all controllers implement the Controller interface
	for _, controller := range controllers {
		assert.Implements(t, (*Controller)(nil), controller)
	}
}

func TestAddHealthEndpoint(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.New()
	addHealthEndpoint(router)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, "healthy", response["status"])
	assert.Contains(t, response, "timestamp")

	// Verify timestamp is a number
	timestamp, ok := response["timestamp"].(float64)
	assert.True(t, ok)
	assert.Greater(t, timestamp, float64(0))
}

func TestRegisterAllRoutes(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create a mock controller
	mockController := &MockController{}
	// Set up expectation
	mockController.On("RegisterRoutes", mock.AnythingOfType("*gin.RouterGroup")).Return()

	controllers := []Controller{mockController}

	router := gin.New()

	routerGroup := router.Group("/")
	registerProtectedRoutes(routerGroup, controllers)

	// Verify that RegisterRoutes was called
	mockController.AssertExpectations(t)
}

// MockController for testing
type MockController struct {
	mock.Mock
}

func (m *MockController) RegisterRoutes(router gin.IRouter) {
	m.Called(router)
}
