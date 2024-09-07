package models

import "gorm.io/gorm"

// User model
type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"` // Email should be unique
	Password string `json:"password"`
}
