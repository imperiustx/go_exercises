package ordertrackingmodel

import "github.com/imperiustx/go_excercises/common"

type OrderTrackingCreate struct {
	common.SQLModelCreate `json:",inline"`
	OrderID               int    `json:"order_id" gorm:"not null"`
	State                 string `json:"state" gorm:"not null"`
}

func (OrderTrackingCreate) TableName() string {
	return OrderTracking{}.TableName()
}
