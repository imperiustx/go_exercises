package restaurantbusiness

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/restaurant/restaurantmodel"
)

// CreateRestaurantStorage create
type CreateRestaurantStorage interface {
	FindRestaurant(
		ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string) (*restaurantmodel.Restaurant, error)
	CreateRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error
}

type createRestaurant struct {
	store CreateRestaurantStorage
}

// NewCreateRestaurantBiz create
func NewCreateRestaurantBiz(store CreateRestaurantStorage) *createRestaurant {
	return &createRestaurant{store: store}
}

func (biz *createRestaurant) CreateNewRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
	restaurant, err := biz.store.FindRestaurant(ctx, map[string]interface{}{"name": data.Name})
	if restaurant != nil {
		return common.ErrEntityExisted(restaurantmodel.EntityName, err)
	}
	
	if err := biz.store.CreateRestaurant(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(restaurantmodel.EntityName, err)
	}

	return nil
}
