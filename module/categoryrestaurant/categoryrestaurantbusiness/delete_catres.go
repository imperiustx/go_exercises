package categoryrestaurantbusiness

import (
	"context"
	"errors"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/categoryrestaurant/categoryrestaurantmodel"
)

// DeleteCategoryRestaurantStorage delete
type DeleteCategoryRestaurantStorage interface {
	FindCategoryRestaurant(ctx context.Context, cid, rid int) (*categoryrestaurantmodel.CategoryRestaurant, error)
	DeleteCategoryRestaurant(cid, rid int) error
}

type deleteCategoryRestaurant struct {
	store DeleteCategoryRestaurantStorage
}

// NewDeleteCategoryRestaurantBiz delete
func NewDeleteCategoryRestaurantBiz(store DeleteCategoryRestaurantStorage) *deleteCategoryRestaurant {
	return &deleteCategoryRestaurant{store: store}
}

func (biz *deleteCategoryRestaurant) DeleteCategoryRestaurant(ctx context.Context, cid, rid int) error {
	categoryrestaurant, err := biz.store.FindCategoryRestaurant(ctx, cid, rid)
	if err != nil {
		return common.ErrCannotGetEntity(categoryrestaurantmodel.EntityName, err)
	}

	if categoryrestaurant.Status == 0 {
		return common.ErrCannotGetEntity(categoryrestaurantmodel.EntityName, errors.New("category restaurant not found"))
	}

	if err := biz.store.DeleteCategoryRestaurant(cid, rid); err != nil {
		return err
	}

	return nil
}
