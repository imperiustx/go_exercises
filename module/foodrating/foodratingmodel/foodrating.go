package foodratingmodel

import (
	"github.com/imperiustx/go_excercises/common"
)

const EntityName = "FoodRating"

// FoodRating model
type FoodRating struct {
	common.SQLModel
	UserID  int     `json:"user_id"`
	FoodID  int     `json:"food_id"`
	Point   float32 `json:"point"`
	Comment string  `json:"comment"`
}

// TableName table name
func (FoodRating) TableName() string {
	return "food_ratings"
}
