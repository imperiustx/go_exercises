package categoryrestaurantmodel

import "github.com/imperiustx/go_excercises/common"

type CategoryRestaurantUpdate struct {
	Icon *common.Image `json:"icon"`
}

func (CategoryRestaurantUpdate) TableName() string {
	return CategoryRestaurant{}.TableName()
}
