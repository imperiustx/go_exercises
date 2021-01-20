package ordermodel

import "github.com/imperiustx/go_excercises/common"

type OrderCreate struct {
	common.SQLModelCreate `json:",inline"`
	UserID                int     `json:"user_id" gorm:"not null"`
	TotalPrice            float64 `json:"total_price" gorm:"not null"`
}

func (OrderCreate) TableName() string {
	return Order{}.TableName()
}
