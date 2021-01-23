package restaurantratingmodel

import (
	"github.com/imperiustx/go_excercises/common"
)

const EntityName = "RestaurantRating"

// RestaurantRating restaurantrating
type RestaurantRating struct {
	common.SQLModel
	UserID       int     `json:"user_id" gorm:"not null"`
	RestaurantID int     `json:"restaurant_id" gorm:"not null"`
	Point        float32 `json:"point" gorm:"not null"`
	Comment      string  `json:"comment"`
}

// TableName table name
func (RestaurantRating) TableName() string {
	return "restaurant_ratings"
}
