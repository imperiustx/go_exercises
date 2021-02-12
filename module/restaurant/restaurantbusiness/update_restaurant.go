package restaurantbusiness

import (
	"context"
	"errors"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/restaurant/restaurantmodel"
)

// UpdateRestaurantStorage update
type UpdateRestaurantStorage interface {
	FindRestaurant(
		ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string) (*restaurantmodel.Restaurant, error)
	UpdateRestaurant(
		ctx context.Context,
		conditions map[string]interface{},
		data *restaurantmodel.RestaurantUpdate) error
}

type updateRestaurant struct {
	store UpdateRestaurantStorage
}

// NewUpdateRestaurantBiz update
func NewUpdateRestaurantBiz(store UpdateRestaurantStorage) *updateRestaurant {
	return &updateRestaurant{store: store}
}

func (biz *updateRestaurant) UpdateRestaurant(
	ctx context.Context,
	conditions map[string]interface{},
	data *restaurantmodel.RestaurantUpdate) error {

	restaurant, err := biz.store.FindRestaurant(ctx, conditions)
	if err != nil {
		return common.ErrCannotGetEntity(restaurantmodel.EntityName, err)
	}

	if restaurant.Status == 0 {
		return common.ErrCannotGetEntity(restaurantmodel.EntityName, errors.New("restaurant not found"))
	}

	if err := biz.store.UpdateRestaurant(ctx, conditions, data); err != nil {
		return common.ErrCannotUpdateEntity(restaurantmodel.EntityName, err)
	}

	return nil
}
