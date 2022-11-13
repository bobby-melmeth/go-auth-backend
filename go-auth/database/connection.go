package database

import (
	"log"
	"os"

	"github.com/bobbyMoonward/go-auth-backend/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

	var DB *gorm.DB

	func Connect(){
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	dsn := os.Getenv("DB_CREDENTIALS")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the database")
	}

	DB = db

	db.AutoMigrate(&models.User{})
}