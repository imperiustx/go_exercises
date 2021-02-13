package restaurantfoodmodel

import (
	"time"
)

const EntityName = "RestaurantFood"

// RestaurantFood model
type RestaurantFood struct {
	RestaurantID int        `json:"restaurant_id"`
	FoodID       int        `json:"food_id"`
	Status       int        `json:"status"`
	CreatedAt    *time.Time `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
}

// TableName table name
func (RestaurantFood) TableName() string {
	return "restaurant_foods"
}