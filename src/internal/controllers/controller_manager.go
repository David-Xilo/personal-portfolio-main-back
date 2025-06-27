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
	"safehouse-main-back/src/internal/controllers/endpoints"
	"safehouse-main-back/src/internal/controllers/security"
	swaggerconfig "safehouse-main-back/src/internal/controllers/swagger"
	"safehouse-main-back/src/internal/database"
)

type Controller interface {
	RegisterRoutes(router *gin.Engine)
}

func SetupRoutes(db database.Database) *gin.Engine {
	controllers := getControllers(db)
	router := createRouter()

	registerAllRoutes(router, controllers)

	swaggerconfig.AddSwaggerEndpoint(router)

	return router
}

func createRouter() *gin.Engine {
	router := gin.Default()

	router.Use(security.SecurityHeadersMiddleware())

	limiter := security.NewIPRateLimiter(rate.Limit(5), 30)
	router.Use(security.RateLimitMiddleware(limiter))

	router.Use(security.GetCors())

	return router
}

func getControllers(db database.Database) []Controller {
	var controllers []Controller

	aboutController := endpoints.NewAboutController(db)
	controllers = append(controllers, aboutController)

	techController := endpoints.NewTechController(db)
	controllers = append(controllers, techController)

	gamesController := endpoints.NewGamesController(db)
	controllers = append(controllers, gamesController)

	financeController := endpoints.NewFinanceController(db)
	controllers = append(controllers, financeController)

	return controllers
}

func registerAllRoutes(router *gin.Engine, controllers []Controller) {
	for _, controller := range controllers {
		controller.RegisterRoutes(router)
	}
}
