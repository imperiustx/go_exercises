package restaurantfoodmodel

import (
	"time"

	"github.com/imperiustx/go_excercises/common"
)

const EntityName = "RestaurantFood"

// RestaurantFood model
type RestaurantFood struct {
	RestaurantID     int         `json:"-" gorm:"restaurant_id"`
	FakeRestaurantID *common.UID `json:"restaurant_id" gorm:"-"`
	FoodID           int         `json:"-" gorm:"food_id"`
	FakeFoodID       *common.UID `json:"food_id" gorm:"-"`
	Status           int         `json:"status"`
	CreatedAt        *time.Time  `json:"created_at"`
	UpdatedAt        *time.Time  `json:"updated_at"`
}

// TableName table name
func (RestaurantFood) TableName() string {
	return "restaurant_foods"
}

func (m *RestaurantFood) GenUID(objectType1, objectType2 int, shardId uint32) {
	rid := common.NewUID(uint32(m.RestaurantID), objectType1, shardId)
	m.FakeRestaurantID = &rid

	fid := common.NewUID(uint32(m.FoodID), objectType2, shardId)
	m.FakeFoodID = &fid
}
