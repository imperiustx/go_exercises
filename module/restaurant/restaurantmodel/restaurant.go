package restaurantmodel

import (
	"github.com/imperiustx/go_excercises/common"
)

const EntityName = "Restaurant"

// Restaurant restaurant
type Restaurant struct {
	common.SQLModel
	OwnerID     int           `json:"owner_id" `
	Name        string        `json:"name" `
	Address     string        `json:"addr" `
	CityID      int           `json:"city_id"`
	Latitude    float64       `json:"latitude"`
	Longitude   float64       `json:"longitude"`
	Logo        *common.Image `json:"logo" `
	ShippingFee int           `json:"shipping_fee" `
}

// TableName table name
func (Restaurant) TableName() string {
	return "restaurants"
}
