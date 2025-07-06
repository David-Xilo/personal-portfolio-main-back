package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	configuration "safehouse-main-back/src/internal/config"
	"strings"
)

func SecurityHeadersMiddleware(config configuration.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		isSwagger := strings.HasPrefix(path, "/swagger/") || path == "/"
		isAPIEndpoint := strings.HasPrefix(path, "/auth/") || strings.HasPrefix(path, "/about/") ||
			strings.HasPrefix(path, "/tech/") || strings.HasPrefix(path, "/games/") ||
			strings.HasPrefix(path, "/finance/") || strings.HasPrefix(path, "/health") ||
			strings.HasPrefix(path, "/internal/")
		isProd := config.Environment == "production"

		if (isProd && !isAPIEndpoint) || (!isProd && !isAPIEndpoint && !isSwagger) {
			errorMsg := fmt.Sprintf("Path not allowed %s", c.Request.URL.Path)
			err := fmt.Errorf(errorMsg)

			slog.Error("SecurityHeadersMiddleware: %v", err)

			c.Error(err)

			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": errorMsg})

			return
		}

		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("Referrer-Policy", "strict-origin-when-cross-origin")
		c.Header("X-Frame-Options", "DENY")
		c.Header("X-XSS-Protection", "1; mode=block")
		c.Header("Cache-Control", "no-store, no-cache, must-revalidate, private")
		c.Header("Pragma", "no-cache")
		c.Header("Expires", "0")

		csp := getCSPPolicy(isSwagger)
		c.Header("Content-Security-Policy", csp)

		c.Next()
	}
}

func getCSPPolicy(isSwagger bool) string {
	if isSwagger {
		return "default-src 'self'; " +
			"script-src 'self' 'unsafe-inline'; " +
			"style-src 'self' 'unsafe-inline'; " +
			"img-src 'self' data:; " +
			"font-src 'self'; " +
			"connect-src 'self'; " +
			"frame-ancestors 'none';"
	}

	return "default-src 'none'; frame-ancestors 'none';"
}
