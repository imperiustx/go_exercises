package restaurantbusiness

import (
	"errors"

	"github.com/imperiustx/go_excercises/module/restaurant/restaurantmodel"
)

// DeleteRestaurantStorage delete
type DeleteRestaurantStorage interface {
	FindRestaurant(id int) (*restaurantmodel.Restaurant, error)
	DeleteRestaurant(id int) error
}

type deleteRestaurant struct {
	store DeleteRestaurantStorage
}

// NewDeleteRestaurantBiz delete
func NewDeleteRestaurantBiz(store DeleteRestaurantStorage) *deleteRestaurant {
	return &deleteRestaurant{store: store}
}

func (biz *deleteRestaurant) DeleteRestaurant(id int) error {
	restaurant, err := biz.store.FindRestaurant(id)
	if err != nil {
		return err
	}

	if restaurant.Status == 0 {
		return errors.New("restaurant note found")
	}

	if err := biz.store.DeleteRestaurant(id); err != nil {
		return err
	}

	return nil
}
