package restaurantlikemodel

type RestaurantLikeCreate struct {
	UserID       string `json:"user_id"`
	RestaurantID string `json:"restaurant_id"`
}

func (RestaurantLikeCreate) TableName() string {
	return RestaurantLike{}.TableName()
}
