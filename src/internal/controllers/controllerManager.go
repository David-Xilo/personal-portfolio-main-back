package controllers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func SetupRoutes(dbClient *gorm.DB) *gin.Engine {

	router := createRouter()

	aboutController := &AboutController{db: dbClient}
	aboutController.RegisterRoutes(router)

	techController := &TechController{db: dbClient}

	gamesController := &GamesController{db: dbClient}
	//financeController := &finance.Controller{DBClient: dbClient}

	// Define the routes and corresponding handler functions
	//aboutHandler := http.StripPrefix("/about", http.HandlerFunc(aboutController.HandleAboutRequest))
	techHandler := http.StripPrefix("/tech", http.HandlerFunc(techController.HandleTechRequest))
	gamesHandler := http.StripPrefix("/games", http.HandlerFunc(gamesController.HandleGamesRequest))
	//financeHandler := http.StripPrefix("/finance", http.HandlerFunc(financeController.HandleFinanceRequest))
	//http.Handle("/about/", addCORSHeaders(aboutHandler))
	http.Handle("/tech/", addCORSHeaders(techHandler))
	http.Handle("/games/", addCORSHeaders(gamesHandler))
	//http.Handle("/finance/", addCORSHeaders(financeHandler))

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

// This needs to be more secure
func addSwagger() {
	http.HandleFunc("/swagger/doc.json", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Hit swagger /swagger/doc.json")
		http.ServeFile(w, r, "./docs/swagger.json")
	})
	http.Handle("/swagger/", httpSwagger.WrapHandler)
}

func addCORSHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if origin := req.Header.Get("Origin"); origin != "" {
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Allow-Methods",
				"POST, GET, PUT, DELETE")
			rw.Header().Set("Access-Control-Allow-Headers",
				"Accept, Accept-Language, Content-Type,")
		}

		// Call the next handler
		next.ServeHTTP(rw, req)
	})
}
