package cartmodel

import (
	"time"

	"github.com/imperiustx/go_excercises/common"
)

const EntityName = "Cart"

// Cart model
type Cart struct {
	UserID     int         `json:"-" gorm:"user_id"`
	FakeUserID *common.UID `json:"user_id" gorm:"-"`
	FoodID     int         `json:"-" gorm:"food_id"`
	FakeFoodID *common.UID `json:"food_id" gorm:"-"`
	Quantity   int         `json:"quantity"`
	Status     int         `json:"status" gorm:"default:1"`
	CreatedAt  *time.Time  `json:"created_at"`
	UpdatedAt  *time.Time  `json:"updated_at"`
}

// TableName table name
func (Cart) TableName() string {
	return "carts"
}

func (m *Cart) GenUID(objectType1, objectType2 int, shardId uint32) {
	uid := common.NewUID(uint32(m.UserID), objectType1, shardId)
	m.FakeUserID = &uid

	rid := common.NewUID(uint32(m.FoodID), objectType2, shardId)
	m.FakeFoodID = &rid
}
