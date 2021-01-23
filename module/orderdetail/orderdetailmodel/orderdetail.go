package orderdetailmodel

import (
	"github.com/imperiustx/go_excercises/common"
)

const EntityName = "OrderDetail"

// OrderDetail model
type OrderDetail struct {
	common.SQLModel
	OrderID    int     `json:"order_id" gorm:"not null"`
	FoodOrigin string  `json:"food_orgin"`
	Price      float64 `json:"price" gorm:"not null"`
	Quantity   int     `json:"quantity" gorm:"not null"`
	Discount   float64 `json:"discount" gorm:"not null"`
}

// TableName table name
func (OrderDetail) TableName() string {
	return "order_details"
}

//TODO: fix food_origin, fix discount: default
