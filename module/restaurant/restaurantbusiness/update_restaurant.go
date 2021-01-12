package restaurantbusiness

// UpdateRestaurantStorage update
type UpdateRestaurantStorage interface {
	UpdateRestaurant(id string, v interface{}) error
}

type updateRestaurant struct {
	store UpdateRestaurantStorage
}

// NewUpdateRestaurantBiz update
func NewUpdateRestaurantBiz(store UpdateRestaurantStorage) *updateRestaurant {
	return &updateRestaurant{store: store}
}

func (biz *updateRestaurant) UpdateRestaurant(id string, v interface{}) error {
	return biz.store.UpdateRestaurant(id, v)
}
