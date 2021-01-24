package restaurantlikebusiness

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/restaurantlike/restaurantlikemodel"
)

// ListRestaurantLikeStorage list
type ListRestaurantLikeStorage interface {
	ListRestaurantLike(ctx context.Context, paging *common.Paging) ([]restaurantlikemodel.RestaurantLike, error)
}

type listRestaurantLike struct {
	store ListRestaurantLikeStorage
}

// NewListRestaurantLikeBiz list
func NewListRestaurantLikeBiz(store ListRestaurantLikeStorage) *listRestaurantLike {
	return &listRestaurantLike{store: store}
}

func (biz *listRestaurantLike) ListAllRestaurantLike(ctx context.Context, paging *common.Paging) ([]restaurantlikemodel.RestaurantLike, error) {
	data, err := biz.store.ListRestaurantLike(ctx, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantlikemodel.EntityName, err)
	}

	return data, nil
}
