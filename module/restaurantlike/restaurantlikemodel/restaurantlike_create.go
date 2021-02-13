package restaurantlikemodel

type RestaurantLikeCreate struct {
	UserID       int `json:"user_id"`
	RestaurantID int `json:"restaurant_id"`
}

func (RestaurantLikeCreate) TableName() string {
	return RestaurantLike{}.TableName()
}
