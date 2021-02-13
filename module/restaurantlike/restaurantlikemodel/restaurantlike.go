package restaurantlikemodel

import "time"

const EntityName = "RestaurantLike"

// RestaurantLike model
type RestaurantLike struct {
	UserID       int        `json:"user_id"`
	RestaurantID int        `json:"restaurant_id"`
	CreatedAt    *time.Time `json:"created_at"`
}

// TableName table name
func (RestaurantLike) TableName() string {
	return "restaurant_likes"
}

// TODO: consider update restaurant like, user can like/unlike
