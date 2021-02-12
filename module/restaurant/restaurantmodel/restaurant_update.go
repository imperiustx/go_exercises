package restaurantmodel

import "github.com/imperiustx/go_excercises/common"

type RestaurantUpdate struct {
	Name        string        `json:"name"`
	Address     string        `json:"addr"`
	Logo        *common.Image `json:"logo"`
	Latitude    float64       `json:"latitude"`
	Longitude   float64       `json:"longitude"`
	ShippingFee int           `json:"shipping_fee"`
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}
