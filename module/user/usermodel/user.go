package usermodel

import (
	"gorm.io/gorm"
)

const EntityName = "User"

// User model
type User struct {
	gorm.Model
	FullName    string `json:"full_name" gorm:"not null"`
	Email       string `json:"email" gorm:"not null"`
	Password    string `json:"password" gorm:"not null"`
	PhoneNumber string `json:"phone_number" gorm:"not null"`
	Status      int    `json:"status" gorm:"not null"`
}

// TableName table name
func (User) TableName() string {
	return "users"
}
