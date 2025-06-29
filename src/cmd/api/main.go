package main

import (
	"context"
	"errors"
	"log/slog"
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
		slog.Info("Server starting", "port", config.Port)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("Failed to start server", "error", err)
			os.Exit(1)
		}
	}()

	shutdownChannel := make(chan os.Signal, 1)
	signal.Notify(shutdownChannel, syscall.SIGINT, syscall.SIGTERM)

	<-shutdownChannel
	slog.Info("Shutting down gracefully")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Error during server shutdown", "error", err)
	} else {
		slog.Info("Server stopped gracefully")
	}

	if err := database.CloseDB(gormDB); err != nil {
		slog.Error("Error closing database connection", "error", err)
	} else {
		slog.Info("Database connection closed successfully")
	}

	slog.Info("Application shutdown complete")
}
