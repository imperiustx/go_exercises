package restaurantratingbusiness

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/restaurantrating/restaurantratingmodel"
)

// ListRestaurantRatingStorage list
type ListRestaurantRatingStorage interface {
	ListRestaurantRating(
		ctx context.Context,
		filter *restaurantratingmodel.Filter,
		paging *common.Paging,
		order *common.OrderSort,
		moreKeys ...string) ([]restaurantratingmodel.RestaurantRating, error)
}

type listRestaurantRating struct {
	store ListRestaurantRatingStorage
}

// NewListRestaurantRatingBiz list
func NewListRestaurantRatingBiz(store ListRestaurantRatingStorage) *listRestaurantRating {
	return &listRestaurantRating{store: store}
}

func (biz *listRestaurantRating) ListAllRestaurantRating(
	ctx context.Context,
	filter *restaurantratingmodel.Filter,
	paging *common.Paging,
	order *common.OrderSort,
	moreKeys ...string) ([]restaurantratingmodel.RestaurantRating, error) {

	data, err := biz.store.ListRestaurantRating(ctx, filter, paging, order)
	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantratingmodel.EntityName, err)
	}

	return data, nil
}
