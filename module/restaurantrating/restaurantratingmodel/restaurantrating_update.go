package restaurantratingmodel

type RestaurantRatingUpdate struct {
	Point   float32 `json:"point"`
	Comment string  `json:"comment"`
}

func (RestaurantRatingUpdate) TableName() string {
	return RestaurantRating{}.TableName()
}
