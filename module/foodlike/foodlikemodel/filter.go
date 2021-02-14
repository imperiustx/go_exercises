package foodlikemodel

type Filter struct {
	UserID int `json:"user_id" form:"user_id"`
	FoodID int `json:"food_id" form:"food_id"`
}
