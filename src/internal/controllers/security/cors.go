package security

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func GetCors() gin.HandlerFunc {
	return cors.New(getCORSConfig())
}

func getCORSConfig() cors.Config {

	allowedHeaders := []string{
		"content-type",
		"referer",
		"sec-ch-ua",
		"sec-ch-ua-mobile",
		"sec-ch-ua-platform",
		"user-agent",
		"x-client-version",
		"origin",
		"accept",
	}

	return cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
		},
		AllowMethods:     []string{"GET", "OPTIONS"},
		AllowHeaders:     allowedHeaders,
		AllowCredentials: true,
	}
}
