package models

import "gorm.io/gorm"

// Restaurant restaurant
type Restaurant struct {
	gorm.Model
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
	PriceRange  string `json:"price_range"`
}
