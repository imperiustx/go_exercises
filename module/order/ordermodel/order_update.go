package ordermodel

type OrderUpdate struct {
	TotalPrice float64 `json:"total_price"`
}

func (OrderUpdate) TableName() string {
	return Order{}.TableName()
}
