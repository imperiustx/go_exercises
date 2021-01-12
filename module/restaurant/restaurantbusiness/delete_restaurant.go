package restaurantbusiness

// DeleteRestaurantStorage delete
type DeleteRestaurantStorage interface {
	DeleteRestaurant(id string) error
}

type deleteRestaurant struct {
	store DeleteRestaurantStorage
}

// NewDeleteRestaurantBiz delete
func NewDeleteRestaurantBiz(store DeleteRestaurantStorage) *deleteRestaurant {
	return &deleteRestaurant{store: store}
}

func (biz *deleteRestaurant) DeleteRestaurant(id string) error {
	return biz.store.DeleteRestaurant(id)
}
