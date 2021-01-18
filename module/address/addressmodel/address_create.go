package addressmodel

import "github.com/imperiustx/go_excercises/common"

type AddressCreate struct {
	common.SQLModelCreate `json:",inline"`
	UserID                int     `json:"user_id"`
	CityID                int     `json:"city_id"`
	Address               string  `json:"address"`
	Latitude              float64 `json:"latitude"`
	Longitude             float64 `json:"longitude"`
}

func (AddressCreate) TableName() string {
	return Address{}.TableName()
}
