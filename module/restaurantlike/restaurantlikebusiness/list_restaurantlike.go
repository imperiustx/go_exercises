package restaurantlikebusiness

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/restaurantlike/restaurantlikemodel"
)

// ListRestaurantLikeStorage list
type ListRestaurantLikeStorage interface {
	ListRestaurantLike(
		ctx context.Context,
		filter *restaurantlikemodel.Filter,
		paging *common.Paging,
		order *common.OrderSort,
		moreKeys ...string) ([]restaurantlikemodel.RestaurantLike, error)
}

type listRestaurantLike struct {
	store ListRestaurantLikeStorage
}

// NewListRestaurantLikeBiz list
func NewListRestaurantLikeBiz(store ListRestaurantLikeStorage) *listRestaurantLike {
	return &listRestaurantLike{store: store}
}

func (biz *listRestaurantLike) ListRestaurantLike(
	ctx context.Context,
	filter *restaurantlikemodel.Filter,
	paging *common.Paging,
	order *common.OrderSort,
	moreKeys ...string) ([]restaurantlikemodel.RestaurantLike, error) {

	data, err := biz.store.ListRestaurantLike(ctx, filter, paging, order)
	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantlikemodel.EntityName, err)
	}

	return data, nil
}
