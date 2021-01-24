package imagemodel

type ImageCreate struct {
	URL    string `json:"url" gorm:"not null"`
	Width  int    `json:"width" gorm:"not null"`
	Height int    `json:"height" gorm:"not null"`
}

func (ImageCreate) TableName() string {
	return Image{}.TableName()
}
