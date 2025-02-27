package main

import (
	"log"
	"os"
	"os/signal"
	"safehouse-main-back/src/internal/controllers"
	"safehouse-main-back/src/internal/database"
	"syscall"
)

func main() {
	gormDB := database.InitDB()

	db := database.NewPostgresDB(gormDB)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-c
		log.Println("Shutting down gracefully...")

		if err := database.CloseDB(gormDB); err != nil {
			log.Printf("Error closing database connection: %v\n", err)
		} else {
			log.Println("Database connection closed successfully")
		}

		os.Exit(0)
	}()

	database.ValidateDBSchema(gormDB)

	router := controllers.SetupRoutes(db)

	port := ":4000"
	err := router.Run(port)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
