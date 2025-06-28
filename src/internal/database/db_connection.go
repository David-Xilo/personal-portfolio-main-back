package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"safehouse-main-back/src/internal/models"
)

func InitDB() *gorm.DB {

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL environment variable not set")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
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
		log.Fatal("Database schema is outdated. Please run the migrations first.")
	}
}
