package foodmodel

import (
	"github.com/imperiustx/go_excercises/common"
)

const EntityName = "Food"

// Food model
type Food struct {
	common.SQLModel
	RestaurantID int            `json:"restaurant_id" gorm:"not null;index"`
	CategoryID   int            `json:"category_id" gorm:"not null;index"`
	Name         string         `json:"name" `
	Description  string         `json:"description"`
	Price        float64        `json:"price" `
	Images       *common.Images `json:"images" `
}

// TableName table name
func (Food) TableName() string {
	return "foods"
}
