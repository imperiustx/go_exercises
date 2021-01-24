package restaurantlikemodel

type RestaurantLikeCreate struct {
	UserID       int `json:"user_id" gorm:"primaryKey;not null"`
	RestaurantID int `json:"restaurant_id" gorm:"primaryKey;index;not null"`
}

func (RestaurantLikeCreate) TableName() string {
	return RestaurantLike{}.TableName()
}
