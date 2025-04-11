package models

import "gorm.io/gorm"

type Meal struct {
	gorm.Model
	MealsName string `gorm:"type:varchar(75);not null"`
	MealsTime string `gorm:"type:varchar(14);not null"`
	// FK - User
	CreatedBy uint `json:"created_by" gorm:"not null"`
	User      User `json:"-" gorm:"foreignKey:CreatedBy;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// FK - Dictionary
	Dictionary Dictionary `json:"-" gorm:"foreignKey:MealsTime;references:DictionaryName;constraint:OnDelete:CASCADE;"`
}
