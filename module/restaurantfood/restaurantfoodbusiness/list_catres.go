package restaurantfoodbusiness

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/restaurantfood/restaurantfoodmodel"
)

// ListRestaurantFoodStorage list
type ListRestaurantFoodStorage interface {
	ListRestaurantFood(
		ctx context.Context,
		paging *common.Paging,
		order *common.OrderSort,
		moreKeys ...string) ([]restaurantfoodmodel.RestaurantFood, error)
}

type listRestaurantFood struct {
	store ListRestaurantFoodStorage
}

// NewListRestaurantFoodBiz list
func NewListRestaurantFoodBiz(store ListRestaurantFoodStorage) *listRestaurantFood {
	return &listRestaurantFood{store: store}
}

func (biz *listRestaurantFood) ListAllRestaurantFood(
	ctx context.Context,
	paging *common.Paging,
	order *common.OrderSort,
	moreKeys ...string) ([]restaurantfoodmodel.RestaurantFood, error) {
		
	data, err := biz.store.ListRestaurantFood(ctx, paging, order)
	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantfoodmodel.EntityName, err)
	}

	return data, nil
}

//TODO: fix order sort