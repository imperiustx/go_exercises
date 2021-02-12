package imagemodel

type ImageCreate struct {
	URL    string `json:"url" `
	Width  int    `json:"width" `
	Height int    `json:"height" `
}

func (ImageCreate) TableName() string {
	return Image{}.TableName()
}
