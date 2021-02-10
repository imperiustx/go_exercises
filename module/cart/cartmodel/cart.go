package cartmodel

import (
	"time"
)

const EntityName = "Cart"

// Cart model
type Cart struct {
	UserID    int     `json:"user_id" gorm:"not null"`
	FoodID    int     `json:"food_id" gorm:"not null"`
	Quantity  int     `json:"quantity" gorm:"not null"`
	Status    int        `json:"status" gorm:"column:status;default:1;"`
	CreatedAt *time.Time `json:"created_at" gorm:"created_at;"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"updated_at;"`
}

// TableName table name
func (Cart) TableName() string {
	return "carts"
}
