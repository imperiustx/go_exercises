package restaurantbusiness

// CreateRestaurantStorage create
type CreateRestaurantStorage interface {
	CreateRestaurant(v interface{}) error
}

type createRestaurant struct {
	store CreateRestaurantStorage
}

// NewCreateRestaurantBiz create
func NewCreateRestaurantBiz(store CreateRestaurantStorage) *createRestaurant {
	return &createRestaurant{store: store}
}

func (biz *createRestaurant) CreateRestaurant(v interface{}) error {
	return biz.store.CreateRestaurant(v)
}
