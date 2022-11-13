package database

import (
	"github.com/bobbyMoonward/go-auth-backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

	var DB *gorm.DB

	func Connect(){
	dsn := "host=localhost user=postgres password=Petroski1 dbname=go-auth port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the database")
	}

	DB = db

	db.AutoMigrate(&models.User{})
}