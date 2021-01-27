package foodmodel

import "github.com/imperiustx/go_excercises/common"

type FoodUpdate struct {
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Price       float64        `json:"price"`
	Images      *common.Images `json:"images"`
}

func (FoodUpdate) TableName() string {
	return Food{}.TableName()
}
