package restaurantmodel

import "gorm.io/gorm"

const EntityName = "Restaurant"

// Restaurant restaurant
type Restaurant struct {
	gorm.Model
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
	PriceRange  string `json:"price_range"`
	Status      int    `json:"status"`
}

// TableName table name
func (Restaurant) TableName() string {
	return "restaurants"
}
