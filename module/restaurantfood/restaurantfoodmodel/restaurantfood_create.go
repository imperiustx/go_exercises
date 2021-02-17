package restaurantfoodmodel

type RestaurantFoodCreate struct {
	RestaurantID string `json:"restaurant_id"`
	FoodID       string `json:"food_id"`
	Status       *int   `json:"status" gorm:"default:1"`
}

func (RestaurantFoodCreate) TableName() string {
	return RestaurantFood{}.TableName()
}
