package foodlikemodel

type FoodLikeCreate struct {
	UserID int `json:"user_id"`
	FoodID int `json:"food_id"`
}

func (FoodLikeCreate) TableName() string {
	return FoodLike{}.TableName()
}
