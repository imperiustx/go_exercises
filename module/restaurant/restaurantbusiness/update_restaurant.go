package restaurantbusiness

import (
	"context"
	"errors"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/restaurant/restaurantmodel"
)

// UpdateRestaurantStorage update
type UpdateRestaurantStorage interface {
	FindRestaurant(ctx context.Context, id int) (*restaurantmodel.Restaurant, error)
	UpdateRestaurant(ctx context.Context, id int, data *restaurantmodel.RestaurantUpdate) error
}

type updateRestaurant struct {
	store UpdateRestaurantStorage
}

// NewUpdateRestaurantBiz update
func NewUpdateRestaurantBiz(store UpdateRestaurantStorage) *updateRestaurant {
	return &updateRestaurant{store: store}
}

func (biz *updateRestaurant) UpdateRestaurant(ctx context.Context, id int, data *restaurantmodel.RestaurantUpdate) error {
	restaurant, err := biz.store.FindRestaurant(ctx, id)
	if err != nil {
		return common.ErrCannotGetEntity(restaurantmodel.EntityName, err)
	}

	if restaurant.Status == 0 {
		return common.ErrCannotGetEntity(restaurantmodel.EntityName, errors.New("restaurant not found"))
	}

	if err := biz.store.UpdateRestaurant(ctx, restaurant.ID, data); err != nil {
		return common.ErrCannotUpdateEntity(restaurantmodel.EntityName, err)
	}

	return nil
}
