package models

import "gorm.io/gorm"

type Social struct {
	gorm.Model
	SocialName        string  `json:"social_name" gorm:"type:varchar(75);not null"`
	SocialPhone       *string `json:"social_phone" gorm:"type:varchar(16)"`
	SocialEmail       *string `json:"social_email" gorm:"type:varchar(255)"`
	SocialGender      string  `json:"social_gender" gorm:"type:varchar(6);not null"`
	SocialAddress     *string `json:"social_address" gorm:"type:varchar(255)"`
	SocialDescription *string `json:"social_description" gorm:"type:varchar(255)"`
	// FK - User
	CreatedBy uint `json:"created_by" gorm:"not null"`
	User      User `json:"-" gorm:"foreignKey:CreatedBy;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// FK - Dictionary
	SocialType uint       `json:"social_type" gorm:"not null"`
	Dictionary Dictionary `gorm:"foreignKey:SocialType;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
