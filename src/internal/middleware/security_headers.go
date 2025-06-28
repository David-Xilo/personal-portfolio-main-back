package middleware

import (
	"github.com/gin-gonic/gin"
	configuration "safehouse-main-back/src/internal/config"
	"strings"
)

func SecurityHeadersMiddleware(config configuration.Config) gin.HandlerFunc {
	return func(c *gin.Context) {

		path := c.Request.URL.Path
		isSwagger := strings.HasPrefix(path, "/swagger/") || path == "/"

		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("Referrer-Policy", "strict-origin-when-cross-origin")
		c.Header("X-Frame-Options", "DENY")
		c.Header("X-XSS-Protection", "1; mode=block")
		c.Header("Cache-Control", "no-store, no-cache, must-revalidate, private")
		c.Header("Pragma", "no-cache")
		c.Header("Expires", "0")

		if isSwagger && config.Environment == "development" {
			c.Header("Content-Security-Policy", "default-src 'self' 'unsafe-inline'; script-src 'self' 'unsafe-inline' 'unsafe-eval'; style-src 'self' 'unsafe-inline';")
		} else {
			c.Header("Content-Security-Policy", "default-src 'none'; frame-ancestors 'none';")
		}

		c.Next()
	}
}
