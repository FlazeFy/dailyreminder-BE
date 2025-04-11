package routes

import (
	"dailyreminder/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetUpRoutes(r *gin.Engine, db *gorm.DB) {
	authController := controllers.NewAuthController(db)
	feedbackController := controllers.NewFeedbackController(db)
	historyController := controllers.NewHistoryController(db)

	api := r.Group("/api/v2")
	{
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
		history := api.Group("/history")
		{
			history.GET("/", historyController.GetAllHistory)
			history.DELETE("/destroy/:id", historyController.HardDeleteHistoryById)
		}
	}
}
