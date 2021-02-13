package restaurantratingmodel

import (
	"github.com/imperiustx/go_excercises/common"
)

const EntityName = "RestaurantRating"

// RestaurantRating restaurantrating
type RestaurantRating struct {
	common.SQLModel
	UserID       int     `json:"user_id"`
	RestaurantID int     `json:"restaurant_id"`
	Point        float32 `json:"point"`
	Comment      string  `json:"comment"`
}

// TableName table name
func (RestaurantRating) TableName() string {
	return "restaurant_ratings"
}
