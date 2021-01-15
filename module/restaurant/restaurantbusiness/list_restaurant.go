package restaurantbusiness

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/restaurant/restaurantmodel"
)

// ListRestaurantStorage list
type ListRestaurantStorage interface {
	ListRestaurant(ctx context.Context, paging *common.Paging) ([]restaurantmodel.Restaurant, error)
}

type listRestaurant struct {
	store ListRestaurantStorage
}

// NewListRestaurantBiz list
func NewListRestaurantBiz(store ListRestaurantStorage) *listRestaurant {
	return &listRestaurant{store: store}
}

func (biz *listRestaurant) ListAllRestaurant(ctx context.Context, paging *common.Paging) ([]restaurantmodel.Restaurant, error) {
	data, err := biz.store.ListRestaurant(ctx, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantmodel.EntityName, err)
	}

	return data, nil
}
