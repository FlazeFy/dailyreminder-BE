package controllers

import (
	"dailyreminder/models"
	"dailyreminder/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type QuestionController struct {
	DB *gorm.DB
}

func NewQuestionController(db *gorm.DB) *QuestionController {
	return &QuestionController{DB: db}
}

// Queries
func (c *QuestionController) GetAllQuestion(ctx *gin.Context) {
	// Models
	var data []models.Question

	// Query
	c.DB.Find(&data)

	// Response
	status := http.StatusNotFound
	var res interface{} = nil

	if len(data) > 0 {
		status = http.StatusOK
		res = data
	}

	ctx.JSON(status, gin.H{
		"data":    res,
		"message": "question fetched",
	})
}

func (c *QuestionController) GetAllUnansweredQuestion(ctx *gin.Context) {
	// Models
	var data []models.Question

	// Query
	if err := c.DB.Where("answer is null").Find(&data).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "unanswered question not found",
		})
		return
	}

	// Response
	status := http.StatusNotFound
	var res interface{} = nil

	if len(data) > 0 {
		status = http.StatusOK
		res = data
	}

	ctx.JSON(status, gin.H{
		"data":    res,
		"message": "unanswered question fetched",
	})
}

// Command
func (c *QuestionController) CreateQuestion(ctx *gin.Context) {
	// Models
	var req models.Question

	// Validate
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request body",
		})
		return
	}

	// Get User ID
	userId, err := utils.GetUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "user not found",
		})
		return
	}

	// Query : Add Question
	question := models.Question{
		Question:  req.Question,
		CreatedBy: userId,
	}
	if err := c.DB.Create(&question).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to create question",
		})
		return
	}

	// Response
	ctx.JSON(http.StatusCreated, gin.H{
		"data":    question,
		"message": "question created",
	})
}

func (c *QuestionController) CreateAnswerQuestion(ctx *gin.Context) {
	// Params
	id := ctx.Param("id")

	// Models
	var req models.AnswerQuestion

	// Validate
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request body",
		})
		return
	}

	if req.Answer == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "answer is required",
		})
		return
	}

	// Check if question exist
	var question models.Question
	if err := c.DB.Where("id = ?", id).First(&question).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "question not found",
		})
		return
	}

	question.Answer = &req.Answer

	// Update Question
	if err := c.DB.Save(&question).Error; err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to update"})
		return
	}

	// Response
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "question answered",
	})
}

func (c *QuestionController) HardDeleteQuestionById(ctx *gin.Context) {
	// Params
	id := ctx.Param("id")

	// Models
	var question models.Question

	// Query
	result := c.DB.Unscoped().First(&question, "id = ?", id)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "question not found",
		})
		return
	}
	c.DB.Unscoped().Delete(&question)

	// Response
	ctx.JSON(http.StatusOK, gin.H{
		"message": "question permanentally deleted",
	})
}
