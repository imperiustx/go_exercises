package foodbusiness

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/food/foodmodel"
)

// GetFoodStorage get
type GetFoodStorage interface {
	FindFood(ctx context.Context, id int) (*foodmodel.Food, error)
}

type getFoodBiz struct {
	store GetFoodStorage
}

// NewGetFoodBiz get
func NewGetFoodBiz(store GetFoodStorage) *getFoodBiz {
	return &getFoodBiz{store: store}
}

func (biz *getFoodBiz) GetFood(ctx context.Context, id int) (*foodmodel.Food, error) {
	food, err := biz.store.FindFood(ctx, id)
	if err != nil {
		return nil, common.ErrCannotGetEntity(foodmodel.EntityName, err)
	}

	return food, nil
}
