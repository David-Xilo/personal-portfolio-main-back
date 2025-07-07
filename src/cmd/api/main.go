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
	"safehouse-main-back/src/internal/secrets"
	"safehouse-main-back/src/internal/security"
	"syscall"
	"time"
)

func main() {
	ctx := context.Background()

	secretProvider, err := secrets.NewSecretProvider(ctx)
	if err != nil {
		slog.Error("Failed to initialize secret manager", "error", err)
		os.Exit(1)
	}

	defer func(secretManager secrets.SecretProvider) {
		err := secretManager.Close()
		if err != nil {
			slog.Error("Failed to Close secret manager", "error", err)
			os.Exit(1)
		}
	}(secretProvider)

	appSecrets, err := secretProvider.LoadAppSecrets(ctx)
	if err != nil {
		slog.Error("Failed to load application secrets", "error", err)
		os.Exit(1)
	}

	gormDB := database.InitDB()
	db := database.NewPostgresDB(gormDB)

	database.ValidateDBSchema(gormDB)

	config := configuration.LoadConfig(appSecrets)

	jwtManager := security.NewJWTManager(config)

	routerSetup := controllers.SetupRoutes(db, config, jwtManager)

	server := &http.Server{
		Addr:         ":" + config.Port,
		Handler:      routerSetup.Router,
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

	routerSetup.RateLimiter.Stop()
	slog.Info("Rate limiter cleanup routine stopped")

	if err := database.CloseDB(gormDB); err != nil {
		slog.Error("Error closing database connection", "error", err)
	} else {
		slog.Info("Database connection closed successfully")
	}

	slog.Info("Application shutdown complete")
}
