package ordermodel

import (
	"github.com/imperiustx/go_excercises/common"
)

const EntityName = "Order"

// Order model
type Order struct {
	common.SQLModel
	UserID     int     `json:"user_id"`
	TotalPrice float64 `json:"total_price"`
	ShipperID  int     `json:"shipper_id"`
}

// TableName table name
func (Order) TableName() string {
	return "orders"
}
