package ordertrackingmodel

import (
	"github.com/imperiustx/go_excercises/common"
)

const EntityName = "OrderTracking"

type OrderTracking struct {
	common.SQLModel
	OrderID int    `json:"order_id" gorm:"not null"`
	State   string `json:"state" gorm:"not null"`
}

func (OrderTracking) TableName() string {
	return "order_trackings"
}
