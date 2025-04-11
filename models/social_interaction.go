package models

import "gorm.io/gorm"

type SocialInteraction struct {
	gorm.Model
	SocialID         string `gorm:"type:uuid;not null"`
	InteractionsMood string `gorm:"type:varchar(36);not null"`
	// FK - User
	CreatedBy uint `json:"created_by" gorm:"not null"`
	User      User `json:"-" gorm:"foreignKey:CreatedBy;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// FK - Dictionary
	Dictionary Dictionary `gorm:"foreignKey:InteractionsMood;references:DictionaryName;constraint:OnDelete:CASCADE;"`
}
