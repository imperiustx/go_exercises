package useraddressmodel

import "github.com/imperiustx/go_excercises/common"

type UserAddressCreate struct {
	common.SQLModelCreate `json:",inline"`
	UserID                int           `json:"user_id"`
	CityID                int           `json:"city_id"`
	Title                 string        `json:"title"`
	Icon                  *common.Image `json:"icon"`
	Addr                  string        `json:"addr"`
	Lat                   float64       `json:"lat"`
	Lng                   float64       `json:"lng"`
}

func (UserAddressCreate) TableName() string {
	return UserAddress{}.TableName()
}
