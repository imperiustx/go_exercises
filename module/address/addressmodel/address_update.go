package addressmodel

type AddressUpdate struct {
	UserID    int     `json:"user_id"`
	CityID    int     `json:"city_id"`
	Address   string  `json:"address"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func (AddressUpdate) TableName() string {
	return Address{}.TableName()
}
