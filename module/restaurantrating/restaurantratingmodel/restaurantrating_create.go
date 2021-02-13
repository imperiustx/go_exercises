package restaurantratingmodel

import "github.com/imperiustx/go_excercises/common"

type RestaurantRatingCreate struct {
	common.SQLModelCreate `json:",inline"`
	UserID                int     `json:"user_id"`
	RestaurantID          int     `json:"restaurant_id"`
	Point                 float32 `json:"point" gorm:"default:0"`
	Comment               string  `json:"comment"`
}

func (RestaurantRatingCreate) TableName() string {
	return RestaurantRating{}.TableName()
}
