package foodmodel

import "github.com/imperiustx/go_excercises/common"

type FoodCreate struct {
	common.SQLModelCreate `json:",inline"`
	Description           string         `json:"description"`
	Price                 float64        `json:"price"`
	Images                *common.Images `json:"logo"`
}

func (FoodCreate) TableName() string {
	return Food{}.TableName()
}
