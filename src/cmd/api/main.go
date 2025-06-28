package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	configuration "safehouse-main-back/src/internal/config"
	"safehouse-main-back/src/internal/controllers"
	"safehouse-main-back/src/internal/database"
	"syscall"
	"time"
)

func main() {
	gormDB := database.InitDB()
	db := database.NewPostgresDB(gormDB)

	database.ValidateDBSchema(gormDB)

	config := configuration.LoadConfig()
	router := controllers.SetupRoutes(db)

	server := &http.Server{
		Addr:         ":" + config.Port,
		Handler:      router,
		ReadTimeout:  config.ReadTimeout,
		WriteTimeout: config.WriteTimeout,
	}

	go func() {
		log.Printf("Server starting on port %s", config.Port)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	shutdownChannel := make(chan os.Signal, 1)
	signal.Notify(shutdownChannel, syscall.SIGINT, syscall.SIGTERM)

	<-shutdownChannel
	log.Println("Shutting down gracefully...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Error during server shutdown: %v", err)
	} else {
		log.Println("Server stopped gracefully")
	}

	if err := database.CloseDB(gormDB); err != nil {
		log.Printf("Error closing database connection: %v", err)
	} else {
		log.Println("Database connection closed successfully")
	}

	log.Println("Application shutdown complete")
}
