package main

import (
	"dailyreminder/config"
	"dailyreminder/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func main() {
	// Load Env
	err := godotenv.Load()

	if err != nil {
		panic("Error loading ENV")
	}

	db := config.ConnectDatabase()
	MigrateAll(db)

	router := gin.Default()
	MigrateAll(db)
	router.Run(":8080")
}

func MigrateAll(db *gorm.DB) {
	db.AutoMigrate(
		&models.User{},
		&models.Admin{},
		&models.Error{},
		&models.Dictionary{},
		&models.History{},
		&models.Social{},
		&models.Alarm{},
		&models.Diary{},
	)
}
