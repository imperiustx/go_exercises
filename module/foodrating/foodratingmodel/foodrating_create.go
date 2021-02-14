package foodratingmodel

import "github.com/imperiustx/go_excercises/common"

type FoodRatingCreate struct {
	common.SQLModelCreate `json:",inline"`
	UserID                int     `json:"user_id"`
	FoodID                int     `json:"food_id"`
	Point                 float32 `json:"point" gorm:"default:0"`
	Comment               string  `json:"comment"`
}

func (FoodRatingCreate) TableName() string {
	return FoodRating{}.TableName()
}
