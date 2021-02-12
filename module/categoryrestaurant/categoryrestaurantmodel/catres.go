package categoryrestaurantmodel

import (
	"time"

	"github.com/imperiustx/go_excercises/common"
)

const EntityName = "CategoryRestaurant"

// CategoryRestaurant model
type CategoryRestaurant struct {
	CategoryID   int           `json:"category_id" gorm:"primaryKey;not null"`
	RestaurantID int           `json:"restaurant_id" gorm:"primaryKey;index;not null"`
	Icon         *common.Image `json:"icon"`
	Status       int           `json:"status" `
	CreatedAt    *time.Time    `json:"created_at"`
	UpdatedAt    *time.Time    `json:"updated_at"`
}

// TableName table name
func (CategoryRestaurant) TableName() string {
	return "category_restaurants"
}
