package routes

import (
	"dailyreminder/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetUpRoutes(r *gin.Engine, db *gorm.DB) {
	feedbackController := controllers.NewFeedbackController(db)

	api := r.Group("/api/v2")
	{
		feedback := api.Group("/feedback")
		{
			feedback.GET("/", feedbackController.GetAllFeedback)
			feedback.POST("/", feedbackController.CreateFeedback)
			feedback.DELETE("/destroy/:id", feedbackController.HardDeleteFeedbackById)
		}
	}
}
