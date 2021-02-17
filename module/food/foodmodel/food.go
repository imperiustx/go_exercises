package foodmodel

import (
	"github.com/imperiustx/go_excercises/common"
)

const EntityName = "Food"

// Food model
type Food struct {
	common.SQLModel
	RestaurantID     int            `json:"-" gorm:"restaurant_id"`
	FakeRestaurantID *common.UID    `json:"restaurant_id" gorm:"-"`
	CategoryID       int            `json:"-" gorm:"category_id"`
	FakeCategoryID   *common.UID    `json:"category_id" gorm:"-"`
	Name             string         `json:"name"`
	Description      string         `json:"description"`
	Price            float64        `json:"price"`
	Images           *common.Images `json:"images"`
}

// TableName table name
func (Food) TableName() string {
	return "foods"
}

func (m *Food) GenUID(objectType1, objectType2, objectType3 int, shardId uint32) {
	fid := common.NewUID(uint32(m.ID), objectType1, shardId)
	m.FakeID = &fid

	rid := common.NewUID(uint32(m.RestaurantID), objectType2, shardId)
	m.FakeRestaurantID = &rid

	cid := common.NewUID(uint32(m.CategoryID), objectType3, shardId)
	m.FakeCategoryID = &cid
}
