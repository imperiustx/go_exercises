package orderdetailmodel

import "github.com/imperiustx/go_excercises/common"

type OrderDetailCreate struct {
	common.SQLModelCreate `json:",inline"`
	OrderID               int     `json:"order_id" `
	FoodOrigin            string  `json:"food_orgin"`
	Price                 float64 `json:"price" `
	Quantity              int     `json:"quantity" `
	Discount              float64 `json:"discount" `
}

func (OrderDetailCreate) TableName() string {
	return OrderDetail{}.TableName()
}

// TODO: fix food_origin, fix discount: default
