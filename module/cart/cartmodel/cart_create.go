package cartmodel

type CartCreate struct {
	UserID   string `json:"user_id"`
	FoodID   string `json:"food_id"`
	Quantity int    `json:"quantity"`
	Status   *int   `json:"status" gorm:"default:1"`
}

func (CartCreate) TableName() string {
	return Cart{}.TableName()
}
