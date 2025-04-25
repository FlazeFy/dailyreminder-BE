package routes

import (
	"dailyreminder/controllers"
	middleware "dailyreminder/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetUpRoutes(r *gin.Engine, db *gorm.DB) {
	authController := controllers.NewAuthController(db)
	feedbackController := controllers.NewFeedbackController(db)
	historyController := controllers.NewHistoryController(db)
	alarmController := controllers.NewAlarmController(db)

	api := r.Group("/api/v2")
	{
		// Public Routes
		auth := api.Group("/auth")
		{
			auth.POST("/register", authController.Register)
			auth.POST("/login", authController.Login)
		}
		feedback := api.Group("/feedback")
		{
			feedback.GET("/", feedbackController.GetAllFeedback)
			feedback.POST("/", feedbackController.CreateFeedback)
			feedback.DELETE("/destroy/:id", feedbackController.HardDeleteFeedbackById)
		}

		// Protected Routes
		protected := api.Group("/")
		protected.Use(middleware.AuthMiddleware())

		alarm := protected.Group("/alarm")
		{
			alarm.GET("/", alarmController.GetAllAlarm)
			alarm.POST("/", alarmController.CreateAlarm)
			alarm.DELETE("/destroy/:id", alarmController.HardDeleteAlarmById)
		}
		history := protected.Group("/history")
		{
			history.GET("/", historyController.GetAllHistory)
			history.DELETE("/destroy/:id", historyController.HardDeleteHistoryById)
		}
	}
}
