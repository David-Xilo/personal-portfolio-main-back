package endpoints

import (
	"net/http"
	"safehouse-main-back/src/internal/security"

	"github.com/gin-gonic/gin"
	configuration "safehouse-main-back/src/internal/config"
)

type AuthController struct {
	config     configuration.Config
	jwtManager *security.JWTManager
}

type AuthRequest struct {
	AuthKey string `json:"auth_key" binding:"required"`
}

type AuthResponse struct {
	Token     string `json:"token"`
	ExpiresIn int    `json:"expires_in"`
}

func NewAuthController(config configuration.Config, jwtManager *security.JWTManager) *AuthController {
	return &AuthController{
		config:     config,
		jwtManager: jwtManager,
	}
}

func (ac *AuthController) RegisterRoutes(router gin.IRouter) {
	router.POST("/auth/token", ac.handleTokenRequest)
}

// @Summary Generate JWT token for frontend authentication
// @Description Authenticate frontend client and return JWT token
// @Tags auth
// @Accept  json
// @Produce  json
// @Param   auth_request body AuthRequest true "Authentication request"
// @Success 200 {object} AuthResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /auth/token [post]
func (ac *AuthController) handleTokenRequest(c *gin.Context) {
	var req AuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request format",
		})
		return
	}

	if req.AuthKey != ac.config.FrontendAuthKey {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid authentication key",
		})
		return
	}

	token, err := ac.jwtManager.GenerateToken()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to generate token",
		})
		return
	}

	c.JSON(http.StatusOK, AuthResponse{
		Token:     token,
		ExpiresIn: ac.config.JWTExpirationMinutes * 60, // Convert to seconds
	})
}
