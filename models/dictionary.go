package models

import (
	"gorm.io/gorm"
)

type (
	Dictionary struct {
		gorm.Model
		DictionaryType string `gorm:"type:varchar(36);not null"`
		DictionaryName string `gorm:"type:varchar(75);unique;not null"`
	}
)
