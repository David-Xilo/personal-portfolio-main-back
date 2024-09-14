package controllers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Controller interface {
	RegisterRoutes(router *gin.Engine)
}

func SetupRoutes(dbClient *gorm.DB) *gin.Engine {
	controllers := getControllers(dbClient)
	router := createRouter()
	registerAllRoutes(router, controllers)
	return router
}

func createRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	return router
}

func getControllers(dbClient *gorm.DB) []Controller {
	var controllers []Controller

	aboutController := &AboutController{db: dbClient}
	controllers = append(controllers, aboutController)

	techController := &TechController{db: dbClient}
	controllers = append(controllers, techController)

	gamesController := &GamesController{db: dbClient}
	controllers = append(controllers, gamesController)

	financeController := &FinanceController{db: dbClient}
	controllers = append(controllers, financeController)

	return controllers
}

func registerAllRoutes(router *gin.Engine, controllers []Controller) {
	for _, controller := range controllers {
		controller.RegisterRoutes(router)
	}
}
