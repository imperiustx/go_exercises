package foodmodel

import "github.com/imperiustx/go_excercises/common"

type FoodCreate struct {
	common.SQLModelCreate `json:",inline"`
	RestaurantID          int            `json:"restaurant_id"`
	CategoryID            int            `json:"category_id"`
	Name                  string         `json:"name"`
	Description           string         `json:"description"`
	Price                 float64        `json:"price"`
	Images                *common.Images `json:"images"`
}

func (FoodCreate) TableName() string {
	return Food{}.TableName()
}
