package cartmodel

import (
	"time"
)

const EntityName = "Cart"

// Cart model
type Cart struct {
	UserID    int        `json:"user_id"`
	FoodID    int        `json:"food_id"`
	Quantity  int        `json:"quantity"`
	Status    int        `json:"status"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

// TableName table name
func (Cart) TableName() string {
	return "carts"
}
