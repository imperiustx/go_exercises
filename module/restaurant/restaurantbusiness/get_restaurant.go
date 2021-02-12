package restaurantbusiness

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/restaurant/restaurantmodel"
)

// GetRestaurantStorage get
type GetRestaurantStorage interface {
	FindRestaurant(
		ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string) (*restaurantmodel.Restaurant, error)
}

type getRestaurant struct {
	store GetRestaurantStorage
}

// NewGetRestaurantBiz get
func NewGetRestaurantBiz(store GetRestaurantStorage) *getRestaurant {
	return &getRestaurant{store: store}
}

func (biz *getRestaurant) GetRestaurant(
	ctx context.Context,
	conditions map[string]interface{}) (*restaurantmodel.Restaurant, error) {
	restaurant, err := biz.store.FindRestaurant(ctx, conditions)
	if err != nil {
		return nil, common.ErrCannotGetEntity(restaurantmodel.EntityName, err)
	}

	return restaurant, nil
}
