package restaurantmodel

import (
	"github.com/imperiustx/go_excercises/common"
)

const EntityName = "Restaurant"

// Restaurant restaurant
type Restaurant struct {
	common.SQLModel
	Name      string        `json:"name" gorm:"not null"`
	Address   string        `json:"address" gorm:"not null"`
	Latitude  float64       `json:"latitude"`
	Longitude float64       `json:"longitude"`
	Logo      *common.Image `json:"logo" gorm:"not null"`
}

// TableName table name
func (Restaurant) TableName() string {
	return "restaurants"
}
