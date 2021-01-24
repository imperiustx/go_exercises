package imagemodel

const EntityName = "Image"

// Image model
type Image struct {
	ID     int    `json:"id" gorm:"primaryKey;not null"`
	URL    string `json:"url" gorm:"not null"`
	Width  int    `json:"width" gorm:"not null"`
	Height int    `json:"height" gorm:"not null"`
}

// TableName table name
func (Image) TableName() string {
	return "images"
}
