package ordertrackingmodel

import "github.com/imperiustx/go_excercises/common"

type OrderTrackingCreate struct {
	common.SQLModelCreate `json:",inline"`
	OrderID               int    `json:"order_id" `
	State                 string `json:"state" `
}

func (OrderTrackingCreate) TableName() string {
	return OrderTracking{}.TableName()
}
