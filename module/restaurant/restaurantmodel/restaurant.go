package restaurantmodel

import (
	"github.com/imperiustx/go_excercises/common"
)

const EntityName = "Restaurant"

// Restaurant restaurant
type Restaurant struct {
	common.SQLModel
	OwnerID     int           `json:"-" gorm:"owner_id"`
	FakeOwnerID *common.UID   `json:"owner_id" gorm:"-"`
	Name        string        `json:"name"`
	Address     string        `json:"addr" gorm:"column:addr;"`
	CityID      int           `json:"city_id"`
	Latitude    float64       `json:"latitude"`
	Longitude   float64       `json:"longitude"`
	Logo        *common.Image `json:"logo"`
	ShippingFee int           `json:"shipping_fee"`
}

// TableName table name
func (Restaurant) TableName() string {
	return "restaurants"
}

func (m *Restaurant) GenUID(objectType1, objectType2 int, shardId uint32) {
	rid := common.NewUID(uint32(m.ID), objectType1, shardId)
	m.FakeID = &rid

	oid := common.NewUID(uint32(m.OwnerID), objectType2, shardId)
	m.FakeOwnerID = &oid
}
