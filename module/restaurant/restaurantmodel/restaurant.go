package restaurantmodel

import (
	"github.com/imperiustx/go_excercises/common"
)

const EntityName = "Restaurant"

// Restaurant restaurant
type Restaurant struct {
	common.SQLModel
	OwnerID     int           `json:"owner_id" gorm:"not null"`
	Name        string        `json:"name" gorm:"not null"`
	Address     string        `json:"address" gorm:"not null"`
	CityID      int           `json:"city_id"`
	Latitude    float64       `json:"latitude"`
	Longitude   float64       `json:"longitude"`
	Logo        *common.Image `json:"logo" gorm:"not null"`
	ShippingFee int           `json:"shipping_fee" gorm:"not null"`
}

// TableName table name
func (Restaurant) TableName() string {
	return "restaurants"
}
