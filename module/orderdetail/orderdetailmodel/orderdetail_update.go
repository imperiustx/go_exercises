package orderdetailmodel

type OrderDetailUpdate struct {
	Price    float64 `json:"price" `
	Quantity int     `json:"quantity" `
}

func (OrderDetailUpdate) TableName() string {
	return OrderDetail{}.TableName()
}
