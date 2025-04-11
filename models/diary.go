package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type (
	Diary struct {
		gorm.Model
		DiaryBody          string         `json:"diary_body" gorm:"type:varchar(1500);not null"`
		DiaryHapinessLevel int            `json:"diary_hapiness_level" gorm:"not null"`
		DiaryAttachment    datatypes.JSON `json:"diary_attachment" gorm:"type:jsonb;null"`
		IsSocialContain    bool           `json:"is_social_contain"`
		// FK - User
		CreatedBy uint `json:"created_by" gorm:"not null"`
		User      User `json:"-" gorm:"foreignKey:CreatedBy;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	}
)
