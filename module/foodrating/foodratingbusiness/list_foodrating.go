package foodratingbusiness

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/foodrating/foodratingmodel"
)

// ListFoodRatingStorage list
type ListFoodRatingStorage interface {
	ListFoodRating(
		ctx context.Context,
		filter *foodratingmodel.Filter,
		paging *common.Paging,
		order *common.OrderSort,
		moreKeys ...string) ([]foodratingmodel.FoodRating, error)
}

type listFoodRating struct {
	store ListFoodRatingStorage
}

// NewListFoodRatingBiz list
func NewListFoodRatingBiz(store ListFoodRatingStorage) *listFoodRating {
	return &listFoodRating{store: store}
}

func (biz *listFoodRating) ListAllFoodRating(
	ctx context.Context,
	filter *foodratingmodel.Filter,
	paging *common.Paging,
	order *common.OrderSort,
	moreKeys ...string) ([]foodratingmodel.FoodRating, error) {

	data, err := biz.store.ListFoodRating(ctx, filter, paging, order)
	if err != nil {
		return nil, common.ErrCannotListEntity(foodratingmodel.EntityName, err)
	}

	return data, nil
}
