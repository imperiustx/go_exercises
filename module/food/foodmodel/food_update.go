package foodmodel

import "github.com/imperiustx/go_excercises/common"

type FoodUpdate struct {
	Description string         `json:"description"`
	Price       float64        `json:"price"`
	Images      *common.Images `json:"logo"`
}

func (FoodUpdate) TableName() string {
	return Food{}.TableName()
}
