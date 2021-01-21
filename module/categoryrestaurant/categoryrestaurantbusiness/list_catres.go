package categoryrestaurantbusiness

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/categoryrestaurant/categoryrestaurantmodel"
)

// ListCategoryRestaurantStorage list
type ListCategoryRestaurantStorage interface {
	ListCategoryRestaurant(ctx context.Context, paging *common.Paging) ([]categoryrestaurantmodel.CategoryRestaurant, error)
}

type listCategoryRestaurant struct {
	store ListCategoryRestaurantStorage
}

// NewListCategoryRestaurantBiz list
func NewListCategoryRestaurantBiz(store ListCategoryRestaurantStorage) *listCategoryRestaurant {
	return &listCategoryRestaurant{store: store}
}

func (biz *listCategoryRestaurant) ListAllCategoryRestaurant(ctx context.Context, paging *common.Paging) ([]categoryrestaurantmodel.CategoryRestaurant, error) {
	data, err := biz.store.ListCategoryRestaurant(ctx, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(categoryrestaurantmodel.EntityName, err)
	}

	return data, nil
}
