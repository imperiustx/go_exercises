package foodlikemodel

import "time"

const EntityName = "FoodLike"

// FoodLike model
type FoodLike struct {
	UserID    int        `json:"user_id"`
	FoodID    int        `json:"food_id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

// TableName table name
func (FoodLike) TableName() string {
	return "food_likes"
}

// TODO: consider update food like, user can like/unlike
