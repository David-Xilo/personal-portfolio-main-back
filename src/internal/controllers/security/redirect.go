package security

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HttpsRedirectMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Railway sets X-Forwarded-Proto header
		if c.Request.Header.Get("X-Forwarded-Proto") == "http" {
			httpsURL := "https://" + c.Request.Host + c.Request.RequestURI
			c.Redirect(http.StatusMovedPermanently, httpsURL)
			c.Abort()
			return
		}
		c.Next()
	}
}
