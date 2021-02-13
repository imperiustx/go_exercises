package restaurantfoodmodel

type RestaurantFoodCreate struct {
	RestaurantID int `json:"restaurant_id"`
	FoodID       int    `json:"food_id"`
	Status       *int   `json:"status"`
}

func (RestaurantFoodCreate) TableName() string {
	return RestaurantFood{}.TableName()
}
