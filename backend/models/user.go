package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"-"`
}

type Appointment struct {
	gorm.Model
	UserID   uint   `json:"user_id"`
	Email    string `json:"email"`
	Date     string `json:"date"`
	Time     string `json:"time"`
	Doctor   string `json:"doctor"`
	Location string `json:"location"`
}

type Prescription struct {
	gorm.Model
	UserID    uint   `json:"user_id"`
	Email     string `json:"email"`
	Medicine  string `json:"medicine"`
	Dosage    string `json:"dosage"`
	Frequency string `json:"frequency"`
}

type MedicalHistory struct {
	gorm.Model
	UserID    uint   `json:"user_id"`
	Email     string `json:"email"`
	Condition string `json:"condition"`
	Notes     string `json:"notes"`
}
