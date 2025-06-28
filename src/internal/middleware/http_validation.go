package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func BasicRequestValidationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.Request.URL.String()) > 1000 {
			c.JSON(http.StatusRequestEntityTooLarge, gin.H{
				"error": "Request too large",
			})
			c.Abort()
			return
		}

		if c.Request.Method != http.MethodGet && c.Request.Method != http.MethodOptions {
			c.JSON(http.StatusMethodNotAllowed, gin.H{
				"error": "Method not allowed",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
