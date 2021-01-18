package foodbusiness

import (
	"context"
	"errors"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/food/foodmodel"
)

// DeleteFoodStorage delete
type DeleteFoodStorage interface {
	FindFood(ctx context.Context, id int) (*foodmodel.Food, error)
	DeleteFood(id int) error
}

type deleteFood struct {
	store DeleteFoodStorage
}

// NewDeleteFoodBiz delete
func NewDeleteFoodBiz(store DeleteFoodStorage) *deleteFood {
	return &deleteFood{store: store}
}

func (biz *deleteFood) DeleteFood(ctx context.Context, id int) error {
	food, err := biz.store.FindFood(ctx, id)
	if err != nil {
		return common.ErrCannotGetEntity(foodmodel.EntityName, err)
	}

	if food.Status == 0 {
		return common.ErrCannotGetEntity(foodmodel.EntityName, errors.New("food not found"))
	}

	if err := biz.store.DeleteFood(id); err != nil {
		return err
	}

	return nil
}
