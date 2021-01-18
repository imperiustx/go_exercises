package foodbusiness

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/food/foodmodel"
)

// CreateFoodStorage create
type CreateFoodStorage interface {
	CreateFood(ctx context.Context, data *foodmodel.FoodCreate) error
}

type createFood struct {
	store CreateFoodStorage
}

// NewCreateFoodBiz create
func NewCreateFoodBiz(store CreateFoodStorage) *createFood {
	return &createFood{store: store}
}

func (biz *createFood) CreateNewFood(ctx context.Context, data *foodmodel.FoodCreate) error {
	if err := biz.store.CreateFood(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(foodmodel.EntityName, err)
	}

	return nil
}
