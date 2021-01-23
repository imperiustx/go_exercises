package ordermodel

type OrderUpdate struct {
	TotalPrice float64 `json:"total_price" gorm:"not null"`
}

func (OrderUpdate) TableName() string {
	return Order{}.TableName()
}
