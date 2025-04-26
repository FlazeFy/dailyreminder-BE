package controllers

import (
	"dailyreminder/models"
	"dailyreminder/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MealController struct {
	DB *gorm.DB
}

func NewMealController(db *gorm.DB) *MealController {
	return &MealController{DB: db}
}

// Queries
func (c *MealController) GetAllMeal(ctx *gin.Context) {
	// Models
	var data []models.Meal

	// Get User ID
	userId, err := utils.GetUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "user not found",
		})
		return
	}

	// Query
	if err := c.DB.Where("created_by = ?", userId).Find(&data).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "meal not found",
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
		"message": "meal fetched",
	})
}

// Command
func (c *MealController) CreateMeal(ctx *gin.Context) {
	// Models
	var req models.Meal

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

	// Query : Add Meal
	meal := models.Meal{
		MealsName: req.MealsName,
		MealsTime: req.MealsTime,
		CreatedBy: userId,
	}
	if err := c.DB.Create(&meal).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "something went wrong",
		})
		return
	}

	// Response
	ctx.JSON(http.StatusCreated, gin.H{
		"data":    meal,
		"message": "meal created",
	})
}

func (c *MealController) HardDeleteMealById(ctx *gin.Context) {
	// Params
	id := ctx.Param("id")

	// Models
	var meal models.Meal

	// Get User ID
	userId, err := utils.GetUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "user not found",
		})
		return
	}

	// Query
	result := c.DB.Unscoped().First(&meal, "id = ? AND created_by = ?", id, userId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "meal not found",
		})
		return
	}
	c.DB.Unscoped().Delete(&meal)

	// Response
	ctx.JSON(http.StatusOK, gin.H{
		"message": "meal permanentally deleted",
	})
}
