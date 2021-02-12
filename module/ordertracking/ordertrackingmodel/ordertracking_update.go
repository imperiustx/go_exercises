package ordertrackingmodel

type OrderTrackingUpdate struct {
	State string `json:"state" `
}

func (OrderTrackingUpdate) TableName() string {
	return OrderTracking{}.TableName()
}
