package models

// Address address
type Address struct {
	Name     string   `json:"name"`
	Location Location `json:"location" gorm:"embedded"`
}

// Location location
type Location struct {
	Latitude   float64 `json:"latitude"`
	Longtitude float64 `json:"longtitude"`
}
