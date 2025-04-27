package models

import "gorm.io/gorm"

type Question struct {
	gorm.Model
	Question string  `gorm:"type:varchar(144);not null"`
	Answer   *string `gorm:"type:varchar(255)"`
	// FK - User
	CreatedBy uint `json:"created_by" gorm:"not null"`
	User      User `json:"-" gorm:"foreignKey:CreatedBy;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
