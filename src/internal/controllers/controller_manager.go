// @title safehouse
// @version 1.0
// @description safehouse documentation for backend
// @termsOfService http://yourterms.com

// @contact.name API Support
// @contact.url http://www.support.com
// @contact.email support@support.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:4000
// @BasePath /
package controllers

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	_ "safehouse-main-back/docs"
	configuration "safehouse-main-back/src/internal/config"
	"safehouse-main-back/src/internal/controllers/endpoints"
	swaggerconfig "safehouse-main-back/src/internal/controllers/swagger"
	"safehouse-main-back/src/internal/database"
	"safehouse-main-back/src/internal/middleware"
)

type Controller interface {
	RegisterRoutes(router *gin.Engine)
}

func SetupRoutes(db database.Database) *gin.Engine {
	config := configuration.LoadConfig()

	controllers := getControllers(db, config)
	router := createRouter(config)

	registerAllRoutes(router, controllers)

	swaggerconfig.AddSwaggerEndpoint(router)

	return router
}

func createRouter(config configuration.Config) *gin.Engine {
	router := gin.Default()

	router.Use(middleware.BasicRequestValidationMiddleware())

	router.Use(middleware.SecurityHeadersMiddleware())

	if config.EnableHTTPSRedirect { // Railway sets this automatically
		router.Use(middleware.HttpsRedirectMiddleware())
	}

	limiter := middleware.NewIPRateLimiter(rate.Limit(5), 30)
	router.Use(middleware.RateLimitMiddleware(limiter))

	router.Use(middleware.GetCors(config))

	return router
}

func getControllers(db database.Database, config configuration.Config) []Controller {
	var controllers []Controller

	aboutController := endpoints.NewAboutController(db, config)
	controllers = append(controllers, aboutController)

	techController := endpoints.NewTechController(db, config)
	controllers = append(controllers, techController)

	gamesController := endpoints.NewGamesController(db, config)
	controllers = append(controllers, gamesController)

	financeController := endpoints.NewFinanceController(db, config)
	controllers = append(controllers, financeController)

	return controllers
}

func registerAllRoutes(router *gin.Engine, controllers []Controller) {
	for _, controller := range controllers {
		controller.RegisterRoutes(router)
	}
}
