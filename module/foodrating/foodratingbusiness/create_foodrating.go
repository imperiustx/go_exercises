package foodratingbusiness

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/foodrating/foodratingmodel"
)

// CreateFoodRatingStorage create
type CreateFoodRatingStorage interface {
	CreateFoodRating(ctx context.Context, data *foodratingmodel.FoodRatingCreate) error
}

type createFoodRating struct {
	store CreateFoodRatingStorage
}

// NewCreateFoodRatingBiz create
func NewCreateFoodRatingBiz(store CreateFoodRatingStorage) *createFoodRating {
	return &createFoodRating{store: store}
}

func (biz *createFoodRating) CreateNewFoodRating(ctx context.Context, data *foodratingmodel.FoodRatingCreate) error {
	if err := biz.store.CreateFoodRating(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(foodratingmodel.EntityName, err)
	}

	return nil
}
