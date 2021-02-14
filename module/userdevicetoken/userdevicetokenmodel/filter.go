package userdevicetokenmodel

type Filter struct {
	UserID int    `json:"user_id" form:"user_id"`
	OS     string `json:"os" form:"os"`
}
