package restaurantmodel

import "github.com/imperiustx/go_excercises/common"

type RestaurantUpdate struct {
	Name    string        `json:"name"`
	Address string        `json:"address"`
	Logo    *common.Image `json:"logo"`
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}
