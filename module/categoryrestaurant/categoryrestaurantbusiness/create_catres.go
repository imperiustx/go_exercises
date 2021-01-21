package categoryrestaurantbusiness

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/categoryrestaurant/categoryrestaurantmodel"
)

// CreateCategoryRestaurantStorage create
type CreateCategoryRestaurantStorage interface {
	CreateCategoryRestaurant(ctx context.Context, data *categoryrestaurantmodel.CategoryRestaurantCreate) error
}

type createCategoryRestaurant struct {
	store CreateCategoryRestaurantStorage
}

// NewCreateCategoryRestaurantBiz create
func NewCreateCategoryRestaurantBiz(store CreateCategoryRestaurantStorage) *createCategoryRestaurant {
	return &createCategoryRestaurant{store: store}
}

func (biz *createCategoryRestaurant) CreateNewCategoryRestaurant(ctx context.Context, data *categoryrestaurantmodel.CategoryRestaurantCreate) error {
	if err := biz.store.CreateCategoryRestaurant(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(categoryrestaurantmodel.EntityName, err)
	}

	return nil
}
