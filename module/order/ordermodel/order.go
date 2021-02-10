package ordermodel

import (
	"github.com/imperiustx/go_excercises/common"
)

const EntityName = "Order"

// Order model
type Order struct {
	common.SQLModel
	UserID     int     `json:"user_id" gorm:"not null"`
	TotalPrice float64 `json:"total_price" gorm:"not null"`
	ShipperID  int     `json:"shipper_id" gorm:"not null"`
}

// TableName table name
func (Order) TableName() string {
	return "orders"
}
