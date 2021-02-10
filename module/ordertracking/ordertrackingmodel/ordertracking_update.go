package ordertrackingmodel

type OrderTrackingUpdate struct {
	State string `json:"state" gorm:"not null"`
}

func (OrderTrackingUpdate) TableName() string {
	return OrderTracking{}.TableName()
}
