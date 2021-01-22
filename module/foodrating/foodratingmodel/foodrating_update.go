package foodratingmodel

type FoodRatingUpdate struct {
	UserID  int     `json:"user_id" gorm:"not null"`
	FoodID  int     `json:"food_id" gorm:"not null"`
	Point   float32 `json:"point" gorm:"default:0;not null"`
	Comment string  `json:"comment"`
}

func (FoodRatingUpdate) TableName() string {
	return FoodRating{}.TableName()
}
