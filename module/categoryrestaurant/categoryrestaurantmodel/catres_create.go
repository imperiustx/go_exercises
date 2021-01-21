package categoryrestaurantmodel

import (
	"github.com/imperiustx/go_excercises/common"
)

type CategoryRestaurantCreate struct {
	common.SQLModelCreate `json:",inline"`
	CategoryID            string `json:"category_id" gorm:"primaryKey;index;"`
	RestaurantID          string `json:"restaurant_id" gorm:"primaryKey;index;"`
	Status                *int   `json:"status" gorm:"default:1;"`
}

func (CategoryRestaurantCreate) TableName() string {
	return CategoryRestaurant{}.TableName()
}
