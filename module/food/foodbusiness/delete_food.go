package foodbusiness

import (
	"context"
	"errors"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/food/foodmodel"
)

// DeleteFoodStorage delete
type DeleteFoodStorage interface {
	FindFood(
		ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string) (*foodmodel.Food, error)
	DeleteFood(
		ctx context.Context,
		conditions map[string]interface{}) error
}

type deleteFood struct {
	store DeleteFoodStorage
}

// NewDeleteFoodBiz delete
func NewDeleteFoodBiz(store DeleteFoodStorage) *deleteFood {
	return &deleteFood{store: store}
}

func (biz *deleteFood) DeleteFood(
	ctx context.Context,
	conditions map[string]interface{}) error {

	food, err := biz.store.FindFood(ctx, conditions)
	if err != nil {
		return common.ErrCannotGetEntity(foodmodel.EntityName, err)
	}

	if food.Status == 0 {
		return common.ErrCannotGetEntity(foodmodel.EntityName, errors.New("food not found"))
	}

	if err := biz.store.DeleteFood(ctx, conditions); err != nil {
		return err
	}

	return nil
}
