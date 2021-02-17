package foodlikemodel

type FoodLikeCreate struct {
	UserID string `json:"user_id"`
	FoodID string `json:"food_id"`
}

func (FoodLikeCreate) TableName() string {
	return FoodLike{}.TableName()
}
