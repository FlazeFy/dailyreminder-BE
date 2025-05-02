package controllers

import (
	"dailyreminder/models"
	"dailyreminder/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SocialController struct {
	DB *gorm.DB
}

func NewSocialController(db *gorm.DB) *SocialController {
	return &SocialController{DB: db}
}

// Queries
func (c *SocialController) GetAllSocial(ctx *gin.Context) {
	// Models
	var data []models.Social

	// Get User ID
	userId, err := utils.GetUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "user not found",
		})
		return
	}

	// Query
	if err := c.DB.Preload("Dictionary").Where("created_by = ?", userId).Find(&data).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "social not found",
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
		"message": "social fetched",
	})
}

// Command
func (c *SocialController) CreateSocial(ctx *gin.Context) {
	// Models
	var req models.Social

	// Validate
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request body",
		})
		return
	}

	// Validate : Social Gender Rules
	allowedTypes := []string{"male", "female"}
	isValidType := false
	for _, t := range allowedTypes {
		if req.SocialGender == t {
			isValidType = true
			break
		}
	}
	if !isValidType {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "dictionary_type must be one of: " + strings.Join(allowedTypes, ", "),
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

	// Query : Add Social
	social := models.Social{
		SocialType:        req.SocialType,
		SocialName:        req.SocialName,
		SocialPhone:       req.SocialPhone,
		SocialEmail:       req.SocialEmail,
		SocialGender:      req.SocialGender,
		SocialAddress:     req.SocialAddress,
		SocialDescription: req.SocialDescription,
		CreatedBy:         userId,
	}
	if err := c.DB.Create(&social).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "something went wrong",
		})
		return
	}

	// Response
	ctx.JSON(http.StatusCreated, gin.H{
		"data":    social,
		"message": "social created",
	})
}

func (c *SocialController) HardDeleteSocialById(ctx *gin.Context) {
	// Params
	id := ctx.Param("id")

	// Models
	var social models.Social

	// Get User ID
	userId, err := utils.GetUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "user not found",
		})
		return
	}

	// Query
	result := c.DB.Unscoped().First(&social, "id = ? AND created_by = ?", id, userId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "social not found",
		})
		return
	}
	c.DB.Unscoped().Delete(&social)

	// Response
	ctx.JSON(http.StatusOK, gin.H{
		"message": "social permanentally deleted",
	})
}
