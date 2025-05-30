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
	dictionaryController := controllers.NewDictionaryController(db)
	mealsController := controllers.NewMealController(db)
	questionController := controllers.NewQuestionController(db)
	socialController := controllers.NewSocialController(db)

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
		dictionary := protected.Group("/dictionary")
		{
			dictionary.GET("/", dictionaryController.GetAllDictionary)
			dictionary.GET("/:dictionary_type", dictionaryController.GetDictionaryByType)
			dictionary.POST("/", dictionaryController.CreateDictionary)
		}
		history := protected.Group("/history")
		{
			history.GET("/", historyController.GetAllHistory)
			history.DELETE("/destroy/:id", historyController.HardDeleteHistoryById)
		}
		social := protected.Group("/social")
		{
			social.GET("/", socialController.GetAllSocial)
			social.POST("/", socialController.CreateSocial)
			social.DELETE("/destroy/:id", socialController.HardDeleteSocialById)
			interaction := social.Group("/interaction")
			{
				interaction.GET("/", socialController.GetAllSocialInteraction)
				interaction.POST("/", socialController.CreateSocialInteraction)
				interaction.DELETE("/destroy/:id", socialController.HardDeleteSocialInteractionById)
			}
		}
		meal := protected.Group("/meal")
		{
			meal.GET("/", mealsController.GetAllMeal)
			meal.POST("/", mealsController.CreateMeal)
			meal.DELETE("/destroy/:id", mealsController.HardDeleteMealById)
		}
		question := protected.Group("/question")
		{
			question.GET("/", questionController.GetAllQuestion)
			question.GET("/unanswered", questionController.GetAllUnansweredQuestion)
			question.POST("/", questionController.CreateQuestion)
			question.PUT("/answer/:id", questionController.CreateAnswerQuestion)
			question.DELETE("/destroy/:id", questionController.HardDeleteQuestionById)
		}
	}
}
