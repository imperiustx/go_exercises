package foodlikemodel

import (
	"time"

	"github.com/imperiustx/go_excercises/common"
)

const EntityName = "FoodLike"

// FoodLike model
type FoodLike struct {
	UserID     int         `json:"-" gorm:"user_id"`
	FakeUserID *common.UID `json:"user_id" gorm:"-"`
	FoodID     int         `json:"-" gorm:"food_id"`
	FakeFoodID *common.UID `json:"food_id" gorm:"-"`
	CreatedAt  *time.Time  `json:"created_at"`
	UpdatedAt  *time.Time  `json:"updated_at"`
}

// TableName table name
func (FoodLike) TableName() string {
	return "food_likes"
}

func (m *FoodLike) GenUID(objectType1, objectType2 int, shardId uint32) {
	uid := common.NewUID(uint32(m.UserID), objectType1, shardId)
	m.FakeUserID = &uid

	rid := common.NewUID(uint32(m.FoodID), objectType2, shardId)
	m.FakeFoodID = &rid
}
