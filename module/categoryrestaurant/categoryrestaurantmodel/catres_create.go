package categoryrestaurantmodel

type CategoryRestaurantCreate struct {
	CategoryID   string `json:"category_id" gorm:"primaryKey"`
	RestaurantID string `json:"restaurant_id" gorm:"primaryKey;index;"`
	Status       *int   `json:"status" gorm:"default:1;"`
}

func (CategoryRestaurantCreate) TableName() string {
	return CategoryRestaurant{}.TableName()
}
