package restaurantbusiness

import "github.com/imperiustx/go_excercises/module/restaurant/restaurantmodel"

// GetRestaurantStorage get
type GetRestaurantStorage interface {
	GetRestaurant(id string) (restaurantmodel.Restaurant, error)
}

type getRestaurant struct {
	store GetRestaurantStorage
}

// NewGetRestaurantBiz get
func NewGetRestaurantBiz(store GetRestaurantStorage) *getRestaurant {
	return &getRestaurant{store: store}
}

func (biz *getRestaurant) GetRestaurant(id string) (restaurantmodel.Restaurant, error) {
	return biz.store.GetRestaurant(id)
}
