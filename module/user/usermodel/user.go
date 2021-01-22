package usermodel

import (
	"github.com/imperiustx/go_excercises/common"
)

const EntityName = "User"

// User model
type User struct {
	common.SQLModel
	Email     string `json:"email" gorm:"not null"`
	Password  string `json:"password" gorm:"not null"`
	FirstName string `json:"first_name" gorm:"not null"`
	LastName  string `json:"last_name" gorm:"not null"`
	Phone     string `json:"phone"`
	Role      string `json:"role" gorm:"default:user;not null"`
	Salt      string
	Avatar    *common.Image `json:"avatar"`
}

// TableName table name
func (User) TableName() string {
	return "users"
}
