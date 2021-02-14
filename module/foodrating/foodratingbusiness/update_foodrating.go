package foodratingbusiness

import (
	"context"
	"errors"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/foodrating/foodratingmodel"
)

// UpdateFoodRatingStorage update
type UpdateFoodRatingStorage interface {
	FindFoodRating(
		ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string) (*foodratingmodel.FoodRating, error)
	UpdateFoodRating(
		ctx context.Context,
		conditions map[string]interface{},
		data *foodratingmodel.FoodRatingUpdate) error
}

type updateFoodRating struct {
	store UpdateFoodRatingStorage
}

// NewUpdateFoodRatingBiz update
func NewUpdateFoodRatingBiz(store UpdateFoodRatingStorage) *updateFoodRating {
	return &updateFoodRating{store: store}
}

func (biz *updateFoodRating) UpdateFoodRating(
	ctx context.Context,
	conditions map[string]interface{},
	data *foodratingmodel.FoodRatingUpdate) error {

	foodrating, err := biz.store.FindFoodRating(ctx, conditions)
	if err != nil {
		return common.ErrCannotGetEntity(foodratingmodel.EntityName, err)
	}

	if foodrating.Status == 0 {
		return common.ErrCannotGetEntity(
			foodratingmodel.EntityName,
			errors.New("foodrating not found"),
		)
	}

	if err := biz.store.UpdateFoodRating(ctx, conditions, data); err != nil {
		return common.ErrCannotUpdateEntity(foodratingmodel.EntityName, err)
	}

	return nil
}
