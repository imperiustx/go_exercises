package restaurantratingmodel

type Filter struct {
	UserID       int     `json:"user_id" form:"user_id"`
	Point        float32 `json:"point" form:"point"`
}
