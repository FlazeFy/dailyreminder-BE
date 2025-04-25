package controllers

import (
	"dailyreminder/models"
	"dailyreminder/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AlarmController struct {
	DB *gorm.DB
}

func NewAlarmController(db *gorm.DB) *AlarmController {
	return &AlarmController{DB: db}
}

// Queries
func (c *AlarmController) GetAllAlarm(ctx *gin.Context) {
	// Models
	var data []models.Alarm

	// Get User ID
	userId, err := utils.GetUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Query
	if err := c.DB.Where("created_by = ?", userId).Find(&data).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
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
		"message": "alarm fetched",
	})
}

// Command
func (c *AlarmController) CreateAlarm(ctx *gin.Context) {
	// Models
	var req models.Alarm

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
			"message": err.Error(),
		})
		return
	}

	// Validate time format
	if err := utils.ValidateAlarmTimeFormat(req.AlarmTime); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Query : Add Alarm
	alarm := models.Alarm{
		AlarmTime:    req.AlarmTime,
		AlarmContext: req.AlarmContext,
		CreatedBy:    userId,
	}
	if err := c.DB.Create(&alarm).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Response
	ctx.JSON(http.StatusCreated, gin.H{
		"data":    alarm,
		"message": "alarm created",
	})
}

func (c *AlarmController) HardDeleteAlarmById(ctx *gin.Context) {
	// Params
	id := ctx.Param("id")

	// Models
	var alarm models.Alarm

	// Get User ID
	userId, err := utils.GetUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Query
	result := c.DB.Unscoped().First(&alarm, "id = ? AND created_by = ?", id, userId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "alarm not found",
		})
		return
	}
	c.DB.Unscoped().Delete(&alarm)

	// Response
	ctx.JSON(http.StatusOK, gin.H{
		"message": "alarm permanentally deleted",
	})
}
