package restaurantlikemodel

import (
	"time"

	"github.com/imperiustx/go_excercises/common"
)

const EntityName = "RestaurantLike"

// RestaurantLike model
type RestaurantLike struct {
	UserID           int         `json:"-" gorm:"user_id"`
	FakeUserID       *common.UID `json:"user_id" gorm:"-"`
	RestaurantID     int         `json:"-" gorm:"restaurant_id"`
	FakeRestaurantID *common.UID `json:"restaurant_id" gorm:"-"`
	CreatedAt        *time.Time  `json:"created_at"`
}

// TableName table name
func (RestaurantLike) TableName() string {
	return "restaurant_likes"
}

func (m *RestaurantLike) GenUID(objectType1, objectType2 int, shardId uint32) {
	uid := common.NewUID(uint32(m.UserID), objectType1, shardId)
	m.FakeUserID = &uid

	rid := common.NewUID(uint32(m.RestaurantID), objectType2, shardId)
	m.FakeRestaurantID = &rid
}
