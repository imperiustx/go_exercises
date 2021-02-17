package restaurantlikemodel

import "github.com/imperiustx/go_excercises/common"

type Filter struct {
	UserID           string      `json:"-" gorm:"user_id"`
	FakeUserID       *common.UID `json:"user_id" form:"user_id" gorm:"-"`
	RestaurantID     string      `json:"-" gorm:"restaurant_id"`
	FakeRestaurantID *common.UID `json:"restaurant_id" form:"restaurant_id" gorm:"-"`
}
