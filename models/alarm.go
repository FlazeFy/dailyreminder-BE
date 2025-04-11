package models

import (
	"gorm.io/gorm"
)

type (
	Alarm struct {
		gorm.Model
		AlarmTime    string `json:"alarm_time" gorm:"type:varchar(6);not null"`
		AlarmContext string `json:"alarm_context" gorm:"type:varchar(255);not null"`
		// FK - User
		CreatedBy uint `json:"created_by" gorm:"not null"`
		User      User `json:"-" gorm:"foreignKey:CreatedBy;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	}
)
