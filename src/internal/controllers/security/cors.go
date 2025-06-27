package security

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	configuration "safehouse-main-back/src/internal/config"
)

func GetCors(config configuration.Config) gin.HandlerFunc {
	return cors.New(getCORSConfig(config))
}

func getCORSConfig(config configuration.Config) cors.Config {

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
			config.FrontendURL,
		},
		AllowMethods:     []string{"GET", "OPTIONS"},
		AllowHeaders:     allowedHeaders,
		AllowCredentials: true,
	}
}
