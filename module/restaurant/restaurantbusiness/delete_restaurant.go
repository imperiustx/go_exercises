package restaurantbusiness

import (
	"context"
	"errors"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/restaurant/restaurantmodel"
)

// DeleteRestaurantStorage delete
type DeleteRestaurantStorage interface {
	FindRestaurant(ctx context.Context, id int) (*restaurantmodel.Restaurant, error)
	DeleteRestaurant(id int) error
}

type deleteRestaurant struct {
	store DeleteRestaurantStorage
}

// NewDeleteRestaurantBiz delete
func NewDeleteRestaurantBiz(store DeleteRestaurantStorage) *deleteRestaurant {
	return &deleteRestaurant{store: store}
}

func (biz *deleteRestaurant) DeleteRestaurant(ctx context.Context, id int) error {
	restaurant, err := biz.store.FindRestaurant(ctx, id)
	if err != nil {
		return common.ErrCannotGetEntity(restaurantmodel.EntityName, err)
	}

	if restaurant.Status == 0 {
		return common.ErrCannotGetEntity(restaurantmodel.EntityName, errors.New("restaurant note found"))
	}

	if err := biz.store.DeleteRestaurant(id); err != nil {
		return err
	}

	return nil
}
