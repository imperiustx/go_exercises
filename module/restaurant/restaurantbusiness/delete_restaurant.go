package restaurantbusiness

import (
	"context"
	"errors"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/restaurant/restaurantmodel"
)

// DeleteRestaurantStorage delete
type DeleteRestaurantStorage interface {
	FindRestaurant(
		ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string) (*restaurantmodel.Restaurant, error)
	DeleteRestaurant(conditions map[string]interface{}) error
}

type deleteRestaurant struct {
	store DeleteRestaurantStorage
}

// NewDeleteRestaurantBiz delete
func NewDeleteRestaurantBiz(store DeleteRestaurantStorage) *deleteRestaurant {
	return &deleteRestaurant{store: store}
}

func (biz *deleteRestaurant) DeleteRestaurant(ctx context.Context, conditions map[string]interface{}) error {
	restaurant, err := biz.store.FindRestaurant(ctx, conditions)
	if err != nil {
		return common.ErrCannotGetEntity(restaurantmodel.EntityName, err)
	}

	if restaurant.Status == 0 {
		return common.ErrCannotGetEntity(restaurantmodel.EntityName, errors.New("restaurant note found"))
	}

	if err := biz.store.DeleteRestaurant(conditions); err != nil {
		return err
	}

	return nil
}
