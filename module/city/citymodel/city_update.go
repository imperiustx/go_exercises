package citymodel

type CityUpdate struct {
	Title string `json:"title"`
}

func (CityUpdate) TableName() string {
	return City{}.TableName()
}
