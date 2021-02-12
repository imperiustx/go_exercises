package useraddressmodel

type Filter struct {
	UserID int `json:"user_id" form:"user_id"`
	CityID int `json:"city_id" form:"city_id"`
}
