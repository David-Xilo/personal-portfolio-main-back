package database

import (
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

	var sslMode string
	if config.IsProduction() {
		sslMode = "required"
	} else {
		sslMode = "disable"
	}

	dbConfig := config.DatabaseConfig
	encodedPassword := url.QueryEscape(dbConfig.DbPassword)

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		dbConfig.DbUser,
		encodedPassword,
		dbConfig.DbHost,
		dbConfig.DbPort,
		dbConfig.DbName,
		sslMode)

	var db *gorm.DB
	var err error
	for i := 0; i < maxRetries; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			slog.Info("Connected to the database successfully")
			break
		}
		slog.Warn("Retrying to connect to the database", "attempt", i+1)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		slog.Error("Failed to connect to the database", "attempts", maxRetries, "error", err)
		os.Exit(1)
	}

	return db
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
