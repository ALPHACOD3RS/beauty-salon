package db

import (
	"log"

	model "github.com/ALPHACOD3RS/Beauty-Salon/internal/models"
	"gorm.io/driver/sqlite" // Import the database driver of your choice
	"gorm.io/gorm"
)


var DB *gorm.DB

func InitDatabase() *gorm.DB{
	
	db, err := gorm.Open(sqlite.Open("beauty.db"), &gorm.Config{})

	if err != nil{
		log.Fatalf("failed to connect db")
	}

	db.AutoMigrate(&model.User{})

	return db

	
}