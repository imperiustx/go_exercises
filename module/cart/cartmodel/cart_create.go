package cartmodel

type CartCreate struct {
	UserID   int  `json:"user_id"`
	FoodID   int  `json:"food_id"`
	Quantity int  `json:"quantity"`
	Status   *int `json:"status"`
}

func (CartCreate) TableName() string {
	return Cart{}.TableName()
}
