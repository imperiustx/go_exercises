package restaurantbusiness

import "github.com/imperiustx/go_excercises/module/restaurant/restaurantmodel"

// ListRestaurantStorage list
type ListRestaurantStorage interface {
	ListRestaurant() ([]restaurantmodel.Restaurant, error)
}

type listRestaurant struct {
	store ListRestaurantStorage
}

// NewListRestaurantBiz list
func NewListRestaurantBiz(store ListRestaurantStorage) *listRestaurant {
	return &listRestaurant{store: store}
}

func (biz *listRestaurant) ListAllRestaurant() ([]restaurantmodel.Restaurant, error) {
	return biz.store.ListRestaurant()
}
