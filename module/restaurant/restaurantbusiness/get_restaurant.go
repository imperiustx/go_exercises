package restaurantbusiness

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/restaurant/restaurantmodel"
)

// GetRestaurantStorage get
type GetRestaurantStorage interface {
	FindRestaurant(ctx context.Context, id int) (*restaurantmodel.Restaurant, error)
}

type getRestaurant struct {
	store GetRestaurantStorage
}

// NewGetRestaurantBiz get
func NewGetRestaurantBiz(store GetRestaurantStorage) *getRestaurant {
	return &getRestaurant{store: store}
}

func (biz *getRestaurant) GetRestaurant(ctx context.Context, id int) (*restaurantmodel.Restaurant, error) {
	restaurant, err := biz.store.FindRestaurant(ctx, id)
	if err != nil {
		return nil, common.ErrCannotGetEntity(restaurantmodel.EntityName, err)
	}

	return restaurant, nil
}
