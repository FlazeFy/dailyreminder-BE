package models

import "gorm.io/gorm"

type UserManual struct {
	gorm.Model
	ManualTitle      string  `gorm:"type:varchar(144);not null"`
	ManualContent    string  `gorm:"type:varchar(1000);not null"`
	ManualAttachment *string `gorm:"type:text"`
	ManualLocation   string  `gorm:"type:text;not null"`
}
