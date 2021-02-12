package useraddressmodel

import (
	"github.com/imperiustx/go_excercises/common"
)

const EntityName = "UserAddress"

// UserAddress model
type UserAddress struct {
	common.SQLModel
	UserID int           `json:"user_id"`
	CityID int           `json:"city_id"`
	Title  string        `json:"title"`
	Icon   *common.Image `json:"icon"`
	Addr   string        `json:"addr"`
	Lat    float64       `json:"lat"`
	Lng    float64       `json:"lng"`
}

// TableName table name
func (UserAddress) TableName() string {
	return "user_addresses"
}
