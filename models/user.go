package models

import (
	"gorm.io/gorm"
)

// User model
type User struct {
	gorm.Model
	FullName    string `json:"full_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
}
