package orderdetailmodel

import "github.com/imperiustx/go_excercises/common"

type OrderDetailCreate struct {
	common.SQLModelCreate `json:",inline"`
	OrderID               int     `json:"order_id" gorm:"not null"`
	FoodOrigin            string  `json:"food_orgin"`
	Price                 float64 `json:"price" gorm:"not null"`
	Quantity              int     `json:"quantity" gorm:"not null"`
	Discount              float64 `json:"discount" gorm:"not null"`
}

func (OrderDetailCreate) TableName() string {
	return OrderDetail{}.TableName()
}

// TODO: fix food_origin, fix discount: default
