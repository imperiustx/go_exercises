package usermodel

type Filter struct {
	Email      string `json:"email" form:"email"`
	Phone      string `json:"phone" form:"phone"`
	FacebookID string `json:"fb_id" form:"fb_id"`
	GoogleID   string `json:"gg_id" form:"gg_id"`
}
