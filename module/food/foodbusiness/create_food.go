package foodbusiness

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/food/foodmodel"
)

// CreateFoodStorage create
type CreateFoodStorage interface {
	FindFood(
		ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string) (*foodmodel.Food, error)
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
	food, err := biz.store.FindFood(ctx, map[string]interface{}{"name": data.Name})
	if food != nil {
		return common.ErrEntityExisted(foodmodel.EntityName, err)
	}

	if err := biz.store.CreateFood(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(foodmodel.EntityName, err)
	}

	return nil
}

// TODO: fix the logic of find food, this one should be find food by name and restaurant
