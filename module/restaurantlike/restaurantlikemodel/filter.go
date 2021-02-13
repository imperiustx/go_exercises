package restaurantlikemodel

type Filter struct {
	UserID       int `json:"user_id" form:"user_id"`
	RestaurantID int `json:"restaurant_id" form:"restaurant_id"`
}
