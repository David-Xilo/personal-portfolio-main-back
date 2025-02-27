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
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	_ "safehouse-main-back/docs"
	"safehouse-main-back/src/internal/database"
	"strings"
)

type Controller interface {
	RegisterRoutes(router *gin.Engine)
}

func SetupRoutes(db database.Database) *gin.Engine {
	controllers := getControllers(db)
	router := createRouter()
	registerAllRoutes(router, controllers)

	// Add the Swagger route
	router.GET("/", func(c *gin.Context) {
		accept := c.Request.Header.Get("Accept")

		// If it looks like a browser request (wants HTML)
		if strings.Contains(accept, "text/html") {
			c.Redirect(http.StatusFound, "/swagger/index.html")
			return
		}

		// Otherwise, treat it as an API request
		c.JSON(http.StatusOK, gin.H{
			"status":        "API is running",
			"documentation": "/swagger/index.html",
			"version":       "1.0.0",
		})
	})
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}

func createRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	return router
}

func getControllers(db database.Database) []Controller {
	var controllers []Controller

	aboutController := &AboutController{db: db}
	controllers = append(controllers, aboutController)

	techController := &TechController{db: db}
	controllers = append(controllers, techController)

	gamesController := &GamesController{db: db}
	controllers = append(controllers, gamesController)

	//financeController := &FinanceController{db: db}
	//controllers = append(controllers, financeController)

	return controllers
}

func registerAllRoutes(router *gin.Engine, controllers []Controller) {
	for _, controller := range controllers {
		controller.RegisterRoutes(router)
	}
}
