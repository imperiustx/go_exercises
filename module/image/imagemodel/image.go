package imagemodel

const EntityName = "Image"

// Image model
type Image struct {
	ID     int    `json:"id" gorm:"primaryKey;not null"`
	URL    string `json:"url" `
	Width  int    `json:"width" `
	Height int    `json:"height" `
}

// TableName table name
func (Image) TableName() string {
	return "images"
}
