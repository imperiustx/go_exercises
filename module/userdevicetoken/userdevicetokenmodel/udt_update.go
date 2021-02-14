package userdevicetokenmodel

type UserDeviceTokenUpdate struct {
	IsProduction int    `json:"is_production"`
	Token        string `json:"token"`
}

func (UserDeviceTokenUpdate) TableName() string {
	return UserDeviceToken{}.TableName()
}
