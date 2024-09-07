package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"healthcare-app/models" // Import models package
	"log"
)

var DB *gorm.DB

func InitDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("healthcare_app.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// Automatically migrate the schema
	DB.AutoMigrate(&models.User{}) // Reference User model from models package
}
