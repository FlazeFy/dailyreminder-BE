package main

import (
	"dailyreminder/config"
	"dailyreminder/models"
	"dailyreminder/routes"

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
	routes.SetUpRoutes(router, db)
	router.Run(":8080")
}

func MigrateAll(db *gorm.DB) {
	db.AutoMigrate(
		&models.Admin{},
		&models.User{},
		&models.Error{},
		&models.Dictionary{},
		&models.History{},
		&models.Social{},
		&models.SocialInteraction{},
		&models.Alarm{},
		&models.Diary{},
		&models.Meal{},
		&models.Feedback{},
		&models.Task{},
		&models.Question{},
		&models.UserManual{},
	)
}
