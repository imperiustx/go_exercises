package foodbusiness

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/food/foodmodel"
)

type ListFoodByCategoryIDStorage interface {
	ListFoodByCategoryID(
		ctx context.Context,
		id int,
		paging *common.Paging,
		order *common.OrderSort,
		moreKeys ...string) ([]foodmodel.Food, error)
}

type listFoodByCategoryID struct {
	store ListFoodByCategoryIDStorage
}

// NewListCategoryBiz list
func NewListFoodByCategoryIDBiz(store ListFoodByCategoryIDStorage) *listFoodByCategoryID {
	return &listFoodByCategoryID{store: store}
}

func (biz *listFoodByCategoryID) ListAllFoodByCategoryID(
	ctx context.Context,
	id int,
	paging *common.Paging,
	order *common.OrderSort,
	moreKeys ...string) ([]foodmodel.Food, error) {

	data, err := biz.store.ListFoodByCategoryID(ctx, id, paging, order)
	if err != nil {
		return nil, common.ErrCannotListEntity(foodmodel.EntityName, err)
	}

	return data, nil
}
