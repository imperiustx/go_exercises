package foodbusiness

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/food/foodmodel"
)

// ListFoodStorage list
type ListFoodStorage interface {
	ListFood(ctx context.Context, paging *common.Paging) ([]foodmodel.Food, error)
}

type listFood struct {
	store ListFoodStorage
}

// NewListFoodBiz list
func NewListFoodBiz(store ListFoodStorage) *listFood {
	return &listFood{store: store}
}

func (biz *listFood) ListAllFood(ctx context.Context, paging *common.Paging) ([]foodmodel.Food, error) {
	data, err := biz.store.ListFood(ctx, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(foodmodel.EntityName, err)
	}

	return data, nil
}
