package models

import "gorm.io/gorm"

type Person struct {
	*gorm.Model
	Name  string `json:"name"  validate:"required"       gorm:"not null;"`
	Email string `json:"email" validate:"required,email" gorm:"unique;not null;"`
}
