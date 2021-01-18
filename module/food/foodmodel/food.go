package foodmodel

import (
	"github.com/imperiustx/go_excercises/common"
)

const EntityName = "Food"

// Food model
type Food struct {
	common.SQLModel
	Name        string         `json:"name" gorm:"not null"`
	Description string         `json:"description"`
	Price       float64        `json:"price" gorm:"not null"`
	Images      *common.Images `json:"logo" gorm:"not null"`
}

// TableName table name
func (Food) TableName() string {
	return "foods"
}
