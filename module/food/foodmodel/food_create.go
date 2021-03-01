package foodmodel

import "github.com/imperiustx/go_excercises/common"

type FoodCreate struct {
	common.SQLModelCreate `json:",inline"`
	RestaurantID          string         `json:"restaurant_id"`
	CategoryID            string         `json:"category_id"`
	Name                  string         `json:"name"`
	Description           string         `json:"description"`
	Price                 float64        `json:"price"`
	Images                *common.Images `json:"-"`
	ImagesID              []int          `json:"images_id" gorm:"-"`
}

func (FoodCreate) TableName() string {
	return Food{}.TableName()
}
