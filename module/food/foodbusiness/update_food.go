package foodbusiness

import (
	"context"
	"errors"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/food/foodmodel"
)

// UpdateFoodStorage update
type UpdateFoodStorage interface {
	FindFood(ctx context.Context, id int) (*foodmodel.Food, error)
	UpdateFood(ctx context.Context, id int, data *foodmodel.FoodUpdate) error
}

type updateFood struct {
	store UpdateFoodStorage
}

// NewUpdateFoodBiz update
func NewUpdateFoodBiz(store UpdateFoodStorage) *updateFood {
	return &updateFood{store: store}
}

func (biz *updateFood) UpdateFood(ctx context.Context, id int, data *foodmodel.FoodUpdate) error {
	food, err := biz.store.FindFood(ctx, id)
	if err != nil {
		return common.ErrCannotGetEntity(foodmodel.EntityName, err)
	}

	if food.Status == 0 {
		return common.ErrCannotGetEntity(foodmodel.EntityName, errors.New("food not found"))
	}

	if err := biz.store.UpdateFood(ctx, food.ID, data); err != nil {
		return common.ErrCannotUpdateEntity(foodmodel.EntityName, err)
	}

	return nil
}
