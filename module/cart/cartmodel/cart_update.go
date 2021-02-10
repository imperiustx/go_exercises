package cartmodel

type CartUpdate struct {
	Quantity int `json:"quantity"`
}

func (CartUpdate) TableName() string {
	return Cart{}.TableName()
}
