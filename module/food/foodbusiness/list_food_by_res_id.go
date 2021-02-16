package foodbusiness

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/food/foodmodel"
)

type ListFoodByRestaurantIDStorage interface {
	ListFoodByRestaurantID(
		ctx context.Context,
		id int,
		paging *common.Paging,
		order *common.OrderSort,
		moreKeys ...string) ([]foodmodel.Food, error)
}

type listFoodByRestaurantID struct {
	store ListFoodByRestaurantIDStorage
}

// NewListRestaurantBiz list
func NewListFoodByRestaurantIDBiz(store ListFoodByRestaurantIDStorage) *listFoodByRestaurantID {
	return &listFoodByRestaurantID{store: store}
}

func (biz *listFoodByRestaurantID) ListAllFoodByRestaurantID(
	ctx context.Context,
	id int,
	paging *common.Paging,
	order *common.OrderSort,
	moreKeys ...string) ([]foodmodel.Food, error) {

	data, err := biz.store.ListFoodByRestaurantID(ctx, id, paging, order)
	if err != nil {
		return nil, common.ErrCannotListEntity(foodmodel.EntityName, err)
	}

	return data, nil
}
