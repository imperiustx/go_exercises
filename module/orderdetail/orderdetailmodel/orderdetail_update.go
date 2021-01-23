package orderdetailmodel

type OrderDetailUpdate struct {
	Price    float64 `json:"price" gorm:"not null"`
	Quantity int     `json:"quantity" gorm:"not null"`
}

func (OrderDetailUpdate) TableName() string {
	return OrderDetail{}.TableName()
}
