package categoryrestaurantbusiness

import (
	"context"
	"errors"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/categoryrestaurant/categoryrestaurantmodel"
)

// UpdateCategoryRestaurantStorage update
type UpdateCategoryRestaurantStorage interface {
	FindCategoryRestaurant(ctx context.Context, cid, rid int) (*categoryrestaurantmodel.CategoryRestaurant, error)
	UpdateCategoryRestaurant(ctx context.Context, cid, rid int, data *categoryrestaurantmodel.CategoryRestaurantUpdate) error
}

type updateCategoryRestaurant struct {
	store UpdateCategoryRestaurantStorage
}

// NewUpdateCategoryRestaurantBiz update
func NewUpdateCategoryRestaurantBiz(store UpdateCategoryRestaurantStorage) *updateCategoryRestaurant {
	return &updateCategoryRestaurant{store: store}
}

func (biz *updateCategoryRestaurant) UpdateCategoryRestaurant(ctx context.Context, cid, rid int, data *categoryrestaurantmodel.CategoryRestaurantUpdate) error {
	categoryrestaurant, err := biz.store.FindCategoryRestaurant(ctx, cid, rid)
	if err != nil {
		return common.ErrCannotGetEntity(categoryrestaurantmodel.EntityName, err)
	}

	if categoryrestaurant.Status == 0 {
		return common.ErrCannotGetEntity(categoryrestaurantmodel.EntityName, errors.New("categoryrestaurant not found"))
	}

	if err := biz.store.UpdateCategoryRestaurant(ctx, categoryrestaurant.CategoryID, categoryrestaurant.RestaurantID, data); err != nil {
		return common.ErrCannotUpdateEntity(categoryrestaurantmodel.EntityName, err)
	}

	return nil
}
