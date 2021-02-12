package restaurantmodel

type Filter struct {
	OwnerID string `json:"ownner_id" form:"owner_id"`
}