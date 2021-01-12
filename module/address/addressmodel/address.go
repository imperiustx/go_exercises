package addressmodel

import (
	"gorm.io/gorm"
)

// Address address
type Address struct {
	gorm.Model
	FullAddress string  `json:"full_address"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
}


// TableName table name
func (Address) TableName() string {
	return "addresses"
}