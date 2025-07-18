package database

import (
	"context"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log/slog"
	"os"
	configuration "safehouse-main-back/src/internal/config"
	"safehouse-main-back/src/internal/models"
	"time"
)

const maxRetries = 15

func InitDB(config configuration.Config) *gorm.DB {

	//dbConfig := config.DatabaseConfig
	//encodedPassword := url.QueryEscape(dbConfig.DbPassword)

	var dsn string
	var err error

	if config.DatabaseConfig.UseIAMAuth {
		slog.Info("Using IAM authentication for database connection")
		dsn = buildIAMDSN(config)
	} else {
		slog.Info("Using password authentication for database connection")
		dsn = config.DatabaseConfig.DbUrl
	}

	slog.Info("Attempting to connect to database", "dsn_pattern", maskPassword(dsn))

	var db *gorm.DB
	for i := 0; i < maxRetries; i++ {
		db, err = attemptConnection(dsn, i+1) // Use your function instead of gorm.Open directly
		if err == nil {
			slog.Info("Connected to the database successfully")
			break
		}
		if i < maxRetries-1 {
			waitTime := time.Duration(i+1) * 2 * time.Second
			slog.Warn("Connection failed, retrying...",
				"attempt", i+1,
				"max_retries", maxRetries,
				"wait_seconds", waitTime.Seconds(),
				"error", err.Error())
			time.Sleep(waitTime)
		}
	}

	if err != nil {
		slog.Error("Failed to connect to the database", "attempts", maxRetries, "error", err)
		os.Exit(1)
	}

	if err := testConnection(db); err != nil {
		slog.Error("Database connection test failed", "error", err)
		os.Exit(1)
	}

	return db
}

func buildIAMDSN(config configuration.Config) string {
	dbConfig := config.DatabaseConfig

	if config.IsProduction() {
		// Cloud Run with Cloud SQL Proxy and IAM auth
		return fmt.Sprintf("postgres://%s@/%s?host=%s&sslmode=require",
			dbConfig.DbUser,
			dbConfig.DbName,
			dbConfig.DbHost)
	} else {
		// Local development with IAM (less common)
		return fmt.Sprintf("postgres://%s@%s:%s/%s?sslmode=disable",
			dbConfig.DbUser,
			dbConfig.DbHost,
			dbConfig.DbPort,
			dbConfig.DbName)
	}
}

func attemptConnection(dsn string, attempt int) (*gorm.DB, error) {
	slog.Info("Starting database connection attempt", "attempt", attempt, "timeout", "30s")

	// Create a context with timeout for the connection attempt
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	config := &gorm.Config{
		Logger: nil, // Disable GORM's default logger to avoid log spam
	}

	slog.Info("Opening GORM connection...")
	db, err := gorm.Open(postgres.Open(dsn), config)
	if err != nil {
		slog.Error("GORM Open failed", "error", err, "attempt", attempt)
		return nil, fmt.Errorf("gorm.Open failed: %w", err)
	}
	slog.Info("GORM connection opened successfully")

	// Get the underlying sql.DB
	slog.Info("Getting underlying SQL DB...")
	sqlDB, err := db.DB()
	if err != nil {
		slog.Error("Failed to get underlying sql.DB", "error", err)
		return nil, fmt.Errorf("failed to get underlying sql.DB: %w", err)
	}

	// Configure connection pool
	slog.Info("Configuring connection pool...")
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// Test the connection with context
	slog.Info("Pinging database...")
	if err := sqlDB.PingContext(ctx); err != nil {
		slog.Error("Database ping failed", "error", err, "attempt", attempt)
		sqlDB.Close()
		return nil, fmt.Errorf("ping failed: %w", err)
	}
	slog.Info("Database ping successful", "attempt", attempt)

	return db, nil
}

func testConnection(db *gorm.DB) error {
	var result int
	if err := db.Raw("SELECT 1").Scan(&result).Error; err != nil {
		return fmt.Errorf("test query failed: %w", err)
	}

	if result != 1 {
		return fmt.Errorf("test query returned unexpected result: %d", result)
	}

	slog.Info("Database connection test passed")
	return nil
}

func maskPassword(dsn string) string {
	if len(dsn) > 50 {
		return dsn[:30] + "***" + dsn[len(dsn)-10:]
	}
	return "***"
}

func CloseDB(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}

func ValidateDBSchema(db *gorm.DB) {
	if !db.Migrator().HasTable(&models.Contacts{}) {
		slog.Error("Database schema is outdated. Please run the migrations first.")
		os.Exit(1)
	}
}
