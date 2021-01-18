package addressmodel

import (
	"github.com/imperiustx/go_excercises/common"
)

const EntityName = "Address"

// Address model
type Address struct {
	common.SQLModel
	UserID    int     `json:"user_id" gorm:"not null"`
	CityID    int     `json:"city_id" gorm:"not null"`
	Address   string  `json:"address" gorm:"not null"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// TableName table name
func (Address) TableName() string {
	return "addresses"
}
