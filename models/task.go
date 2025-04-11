package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	TaskStatus     string  `gorm:"type:varchar(14);not null"`
	TaskPriority   string  `gorm:"type:varchar(14);not null"`
	TaskTitle      string  `gorm:"type:varchar(75);not null"`
	TaskDesc       *string `gorm:"type:varchar(500)"`
	TaskAttachment *string `gorm:"type:text"`
	TaskAchievment *string `gorm:"type:varchar(144)"`
	TaskDueDate    *time.Time
	TaskCheckpoint *string `gorm:"type:text"`
	// FK - User
	CreatedBy uint `json:"created_by" gorm:"not null"`
	User      User `json:"-" gorm:"foreignKey:CreatedBy;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// FK - Dictionary
	Type     Dictionary `json:"-" gorm:"foreignKey:TaskStatus;references:DictionaryName;constraint:OnDelete:CASCADE;"`
	Priority Dictionary `json:"-" gorm:"foreignKey:TaskPriority;references:DictionaryName;constraint:OnDelete:CASCADE;"`
}
