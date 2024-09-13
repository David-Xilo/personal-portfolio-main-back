package main

import (
	"log"
	"os"
	"os/signal"
	"safehouse-main-back/src/internal/controllers"
	"safehouse-main-back/src/internal/database"
	"syscall"
)

// @title Safehouse API
// @version 1.0
// @description API for Safehouse application.
// @host localhost:4000
// @BasePath /api/v1
func main() {
	db := database.InitDB()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-c
		log.Println("Shutting down gracefully...")

		if err := database.CloseDB(db); err != nil {
			log.Printf("Error closing database connection: %v\n", err)
		} else {
			log.Println("Database connection closed successfully")
		}

		os.Exit(0)
	}()

	database.ValidateDBSchema(db)

	router := controllers.SetupRoutes(db)

	port := ":4000"
	router.Run(port)
}
