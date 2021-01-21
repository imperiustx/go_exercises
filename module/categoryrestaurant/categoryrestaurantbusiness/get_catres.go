package categoryrestaurantbusiness

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/categoryrestaurant/categoryrestaurantmodel"
)

// GetCategoryRestaurantStorage get
type GetCategoryRestaurantStorage interface {
	FindCategoryRestaurant(ctx context.Context, cid, rid int) (*categoryrestaurantmodel.CategoryRestaurant, error)
}

type getCategoryRestaurantBiz struct {
	store GetCategoryRestaurantStorage
}

// NewGetCategoryRestaurantBiz get
func NewGetCategoryRestaurantBiz(store GetCategoryRestaurantStorage) *getCategoryRestaurantBiz {
	return &getCategoryRestaurantBiz{store: store}
}

func (biz *getCategoryRestaurantBiz) GetCategoryRestaurant(ctx context.Context, cid, rid int) (*categoryrestaurantmodel.CategoryRestaurant, error) {
	categoryrestaurant, err := biz.store.FindCategoryRestaurant(ctx, cid, rid)
	if err != nil {
		return nil, common.ErrCannotGetEntity(categoryrestaurantmodel.EntityName, err)
	}

	return categoryrestaurant, nil
}
