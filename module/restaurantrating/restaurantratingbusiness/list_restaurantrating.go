package restaurantratingbusiness

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/restaurantrating/restaurantratingmodel"
)

// ListRestaurantRatingStorage list
type ListRestaurantRatingStorage interface {
	ListRestaurantRating(ctx context.Context, paging *common.Paging) ([]restaurantratingmodel.RestaurantRating, error)
}

type listRestaurantRating struct {
	store ListRestaurantRatingStorage
}

// NewListRestaurantRatingBiz list
func NewListRestaurantRatingBiz(store ListRestaurantRatingStorage) *listRestaurantRating {
	return &listRestaurantRating{store: store}
}

func (biz *listRestaurantRating) ListAllRestaurantRating(ctx context.Context, paging *common.Paging) ([]restaurantratingmodel.RestaurantRating, error) {
	data, err := biz.store.ListRestaurantRating(ctx, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantratingmodel.EntityName, err)
	}

	return data, nil
}
