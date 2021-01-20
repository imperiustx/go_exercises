package categorymodel

import (
	"github.com/imperiustx/go_excercises/common"
)

const EntityName = "Category"

// Category model
type Category struct {
	common.SQLModel
	Name        string        `json:"name" gorm:"not null"`
	Description string        `json:"description"`
	Icon        *common.Image `json:"icon"`
}

// TableName table name
func (Category) TableName() string {
	return "categories"
}
