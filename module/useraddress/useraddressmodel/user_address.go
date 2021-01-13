package useraddressmodel

import (
	"gorm.io/gorm"
)

// UserAddress address
type UserAddress struct {
	gorm.Model
	UserID    int `json:"user_id"`
	AddressID int `json:"address_id"`
}

// TableName table name
func (UserAddress) TableName() string {
	return "user_addresses"
}
