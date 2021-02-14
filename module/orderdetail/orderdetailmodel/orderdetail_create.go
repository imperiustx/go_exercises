package orderdetailmodel

import "github.com/imperiustx/go_excercises/common"

type OrderDetailCreate struct {
	common.SQLModelCreate `json:",inline"`
	OrderID               int           `json:"order_id"`
	FoodOrigin            *common.Image `json:"food_origin"`
	Price                 float64       `json:"price"`
	Quantity              int           `json:"quantity"`
	Discount              float64       `json:"discount" gorm:"default:0"`
}

func (OrderDetailCreate) TableName() string {
	return OrderDetail{}.TableName()
}

// TODO: fix food_origin, fix discount: default
