package foodmodel

type Filter struct {
	RestaurantID int     `json:"restaurant_id" form:"restaurant_id"`
	CategoryID   int     `json:"category_id" form:"category_id"`
	Price        float64 `json:"price" form:"price"`
}
