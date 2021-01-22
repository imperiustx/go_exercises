package foodlikemodel

type FoodLikeCreate struct {
	UserID int `json:"user_id" gorm:"primaryKey;not null"`
	FoodID int `json:"food_id" gorm:"primaryKey;index;not null"`
}

func (FoodLikeCreate) TableName() string {
	return FoodLike{}.TableName()
}
