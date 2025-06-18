package db

import (
	"go_library/internal/infrastructure/db/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func InitDB() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=postgres dbname=calculation port=5433 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	if err := db.AutoMigrate(&models.Author{}, &models.Book{}, &models.Genre{}, &models.User{}); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
	return db, nil
}
