package models

import "gorm.io/gorm"

type SocialInteraction struct {
	gorm.Model
	InteractionsMood string `json:"interactions_mood" gorm:"type:varchar(36);not null"`
	// FK - User
	CreatedBy uint `json:"created_by" gorm:"not null"`
	User      User `json:"-" gorm:"foreignKey:CreatedBy;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// FK - Dictionary
	Dictionary Dictionary `gorm:"foreignKey:InteractionsMood;references:DictionaryName;constraint:OnDelete:CASCADE;"`
	// FK - Social
	SocialID uint   `json:"social_id" gorm:"not null"`
	Social   Social `gorm:"foreignKey:SocialID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
