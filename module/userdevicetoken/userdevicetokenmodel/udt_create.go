package userdevicetokenmodel

import "github.com/imperiustx/go_excercises/common"

type UserDeviceTokenCreate struct {
	common.SQLModelCreate `json:",inline"`
	UserID                int    `json:"user_id"`
	IsProduction          *int   `json:"is_production" gorm:"default:0"`
	OS                    string `json:"os" gorm:"default:ios"`
	Token                 string `json:"token"`
}

func (UserDeviceTokenCreate) TableName() string {
	return UserDeviceToken{}.TableName()
}
