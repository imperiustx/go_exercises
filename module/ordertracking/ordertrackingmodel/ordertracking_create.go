package ordertrackingmodel

import "github.com/imperiustx/go_excercises/common"

type OrderTrackingCreate struct {
	common.SQLModelCreate `json:",inline"`
	OrderID               int    `json:"order_id"`
	State                 string `json:"state" gorm:"default:waiting_for_shipper"`
}

func (OrderTrackingCreate) TableName() string {
	return OrderTracking{}.TableName()
}
