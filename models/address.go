package models

import "gorm.io/gorm"

// Address address
type Address struct {
	gorm.Model
	FullAddress string  `json:"full_address"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longtitude"`
}
