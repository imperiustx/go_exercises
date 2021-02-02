package usermodel

type Filter struct {
	Email      string `json:"email" form:"email"`
	Phone      string `json:"phone" form:"phone"`
}
