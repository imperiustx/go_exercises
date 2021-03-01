package orderdetailmodel

import (
	"github.com/imperiustx/go_excercises/common"
)

const EntityName = "OrderDetail"

// OrderDetail model
type OrderDetail struct {
	common.SQLModel
	OrderID     int           `json:"-" gorm:"order_id"`
	FakeOrderID *common.UID   `json:"order_id" gorm:"-"`
	FoodOrigin  *common.Image `json:"food_origin"` //FIXME: fix food_orgin data type
	Price       float64       `json:"price"`
	Quantity    int           `json:"quantity"`
	Discount    float64       `json:"discount"`
}

// TableName table name
func (OrderDetail) TableName() string {
	return "order_details"
}

func (m *OrderDetail) GenUID(objectType1, objectType2 int, shardId uint32) {
	odid := common.NewUID(uint32(m.ID), objectType1, shardId)
	m.FakeID = &odid

	oid := common.NewUID(uint32(m.OrderID), objectType2, shardId)
	m.FakeOrderID = &oid
}

//TODO: fix food_origin, fix discount: default
