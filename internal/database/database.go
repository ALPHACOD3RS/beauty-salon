package database

import (
	"log"

	"github.com/ALPHACOD3RS/Beauty-Salon/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("beauty.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect db: %v", err)
	}

	db.AutoMigrate(&models.User{}, &models.Payment{}, &models.Appointment{}, &models.Service{})

	DB = db 

	return db
}