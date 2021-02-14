package userdevicetokenmodel

import (
	"github.com/imperiustx/go_excercises/common"
)

const EntityName = "UserDeviceToken"

// UserDeviceToken model
type UserDeviceToken struct {
	common.SQLModel
	UserID       int    `json:"user_id"`
	IsProduction int    `json:"is_production"`
	OS           string `json:"os" `
	Token        string `json:"token"`
}

// TableName table name
func (UserDeviceToken) TableName() string {
	return "user_device_tokens"
}
