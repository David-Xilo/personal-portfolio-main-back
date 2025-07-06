package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log/slog"
	"os"
	"safehouse-main-back/src/internal/models"
	"time"
)

const maxRetries = 15

func InitDB() *gorm.DB {

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		slog.Error("DATABASE_URL environment variable not set")
		os.Exit(1)
	}

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
		slog.Error("Failed to connect to the database", "attempts", maxRetries)
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
