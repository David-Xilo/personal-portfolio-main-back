package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"safehouse-main-back/src/internal/models"
	"time"
)

func InitDB() *gorm.DB {

	maxRetries := 15 // Retry for 30 seconds (with a 2-second interval)
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL environment variable not set")
	}

	var db *gorm.DB
	var err error
	for i := 0; i < maxRetries; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			log.Println("Connected to the database successfully!")
			break
		}
		log.Printf("Retrying to connect to the database... attempt %d, error: %v\n", i+1, err)
		time.Sleep(2 * time.Second)
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
