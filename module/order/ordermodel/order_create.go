package ordermodel

import "github.com/imperiustx/go_excercises/common"

type OrderCreate struct {
	common.SQLModelCreate `json:",inline"`
	UserID                int     `json:"user_id" `
	TotalPrice            float64 `json:"total_price" `
	ShipperID             int     `json:"shipper_id" `
}

func (OrderCreate) TableName() string {
	return Order{}.TableName()
}
