package foodratingbusiness

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/foodrating/foodratingmodel"
)

// ListFoodRatingStorage list
type ListFoodRatingStorage interface {
	ListFoodRating(ctx context.Context, paging *common.Paging) ([]foodratingmodel.FoodRating, error)
}

type listFoodRating struct {
	store ListFoodRatingStorage
}

// NewListFoodRatingBiz list
func NewListFoodRatingBiz(store ListFoodRatingStorage) *listFoodRating {
	return &listFoodRating{store: store}
}

func (biz *listFoodRating) ListAllFoodRating(ctx context.Context, paging *common.Paging) ([]foodratingmodel.FoodRating, error) {
	data, err := biz.store.ListFoodRating(ctx, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(foodratingmodel.EntityName, err)
	}

	return data, nil
}
