package foodlikebusiness

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/foodlike/foodlikemodel"
)

// ListFoodLikeStorage list
type ListFoodLikeStorage interface {
	ListFoodLike(ctx context.Context, paging *common.Paging) ([]foodlikemodel.FoodLike, error)
}

type listFoodLike struct {
	store ListFoodLikeStorage
}

// NewListFoodLikeBiz list
func NewListFoodLikeBiz(store ListFoodLikeStorage) *listFoodLike {
	return &listFoodLike{store: store}
}

func (biz *listFoodLike) ListAllFoodLike(ctx context.Context, paging *common.Paging) ([]foodlikemodel.FoodLike, error) {
	data, err := biz.store.ListFoodLike(ctx, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(foodlikemodel.EntityName, err)
	}

	return data, nil
}
