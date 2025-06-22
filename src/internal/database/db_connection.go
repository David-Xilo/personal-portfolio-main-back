package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"safehouse-main-back/src/internal/models"
	"time"
)

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Dbname   string
	Sslmode  string
}

var config = DBConfig{
	Host:     "safehouse-db-container",
	Port:     5432,
	User:     "safehouse-main-user",
	Password: "mypassword",
	Dbname:   "safehouse-main-db",
	Sslmode:  "disable",
}

func InitDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.Dbname, config.Sslmode)

	var db *gorm.DB
	var err error
	maxRetries := 15 // Retry for 30 seconds (with a 2-second interval)

	for i := 0; i < maxRetries; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			log.Println("Connected to the database successfully!")
			break
		}
		log.Printf("Retrying to connect to the database... attempt %d, error: %v\n", i+1, err)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatalf("Failed to connect to the database after 30 seconds: %v", err)
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
