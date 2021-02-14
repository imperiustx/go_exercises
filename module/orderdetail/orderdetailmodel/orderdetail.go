package orderdetailmodel

import (
	"github.com/imperiustx/go_excercises/common"
)

const EntityName = "OrderDetail"

// OrderDetail model
type OrderDetail struct {
	common.SQLModel
	OrderID    int           `json:"order_id"`
	FoodOrigin *common.Image `json:"food_origin"` //FIXME: fix food_orgin data type
	Price      float64       `json:"price"`
	Quantity   int           `json:"quantity"`
	Discount   float64       `json:"discount"`
}

// TableName table name
func (OrderDetail) TableName() string {
	return "order_details"
}

//TODO: fix food_origin, fix discount: default
