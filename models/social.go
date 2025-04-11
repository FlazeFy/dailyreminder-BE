package models

import "gorm.io/gorm"

type Social struct {
	gorm.Model
	SocialType        string  `gorm:"type:varchar(36);not null"`
	SocialName        string  `gorm:"type:varchar(75);not null"`
	SocialPhone       *string `gorm:"type:varchar(16)"`
	SocialEmail       *string `gorm:"type:varchar(255)"`
	SocialGender      string  `gorm:"type:varchar(6);not null"`
	SocialAddress     *string `gorm:"type:varchar(255)"`
	SocialDescription *string `gorm:"type:varchar(255)"`
	// FK - User
	CreatedBy uint `json:"created_by" gorm:"not null"`
	User      User `json:"-" gorm:"foreignKey:CreatedBy;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
