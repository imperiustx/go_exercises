package restaurantmodel

import "github.com/imperiustx/go_excercises/common"

type RestaurantCreate struct {
	common.SQLModelCreate `json:",inline"`
	Name                  string        `json:"name"`
	Address               string        `json:"address"`
	Latitude              float64       `json:"latitude"`
	Longitude             float64       `json:"longitude"`
	Logo                  *common.Image `json:"logo"`
}

func (RestaurantCreate) TableName() string {
	return Restaurant{}.TableName()
}
