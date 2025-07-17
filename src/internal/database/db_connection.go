package database

import (
	"context"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log/slog"
	"net/url"
	"os"
	configuration "safehouse-main-back/src/internal/config"
	"safehouse-main-back/src/internal/models"
	"time"
)

const maxRetries = 15

func InitDB(config configuration.Config) *gorm.DB {

	//dbConfig := config.DatabaseConfig
	//encodedPassword := url.QueryEscape(dbConfig.DbPassword)

	useIAMAuth := os.Getenv("USE_IAM_DB_AUTH") == "true"

	var dsn string
	var err error

	if useIAMAuth {
		slog.Info("Using IAM authentication for database connection")
		dsn = buildIAMDSN(config)
	} else {
		slog.Info("Using password authentication for database connection")
		dsn, err = buildPasswordDSN(config)
		if err != nil {
			slog.Error("Failed to build password DSN", "error", err)
			os.Exit(1)
		}
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
			slog.Warn("Retrying to connect to the database", "attempt", i+1, "wait_seconds", waitTime.Seconds(), "error", err.Error())
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

func buildPasswordDSN(config configuration.Config) (string, error) {

	dbConfig := config.DatabaseConfig
	userInfo := url.UserPassword(dbConfig.DbUser, dbConfig.DbPassword)

	if config.IsProduction() {
		// Cloud Run with Cloud SQL Proxy
		return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
			userInfo.String(),
			dbConfig.DbHost,
			dbConfig.DbPort,
			dbConfig.DbName), nil
	} else {
		// Local development
		return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
			userInfo.String(),
			dbConfig.DbHost,
			dbConfig.DbPort,
			dbConfig.DbName), nil
	}
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
	slog.Debug("Attempting database connection", "attempt", attempt)

	// Create a context with timeout for the connection attempt
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	config := &gorm.Config{
		Logger: nil, // Disable GORM's default logger to avoid log spam
	}

	db, err := gorm.Open(postgres.Open(dsn), config)
	if err != nil {
		return nil, fmt.Errorf("gorm.Open failed: %w", err)
	}

	// Get the underlying sql.DB
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get underlying sql.DB: %w", err)
	}

	// Configure connection pool
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// Test the connection with context
	if err := sqlDB.PingContext(ctx); err != nil {
		sqlDB.Close()
		return nil, fmt.Errorf("ping failed: %w", err)
	}

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
