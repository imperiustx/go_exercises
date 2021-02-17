package foodlikemodel

import "github.com/imperiustx/go_excercises/common"

type Filter struct {
	UserID     string      `json:"-" gorm:"user_id"`
	FakeUserID *common.UID `json:"user_id" form:"user_id" gorm:"-"`
	FoodID     string      `json:"-" form:"food_id"`
	FakeFoodID *common.UID `json:"food_id" form:"food_id" gorm:"-"`
}
