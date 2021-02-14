package foodratingbusiness

import (
	"context"
	"errors"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/foodrating/foodratingmodel"
)

// DeleteFoodRatingStorage delete
type DeleteFoodRatingStorage interface {
	FindFoodRating(
		ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string) (*foodratingmodel.FoodRating, error)
	DeleteFoodRating(
		ctx context.Context,
		conditions map[string]interface{}) error
}

type deleteFoodRating struct {
	store DeleteFoodRatingStorage
}

// NewDeleteFoodRatingBiz delete
func NewDeleteFoodRatingBiz(store DeleteFoodRatingStorage) *deleteFoodRating {
	return &deleteFoodRating{store: store}
}

func (biz *deleteFoodRating) DeleteFoodRating(
	ctx context.Context,
	conditions map[string]interface{}) error {

	foodrating, err := biz.store.FindFoodRating(ctx, conditions)
	if err != nil {
		return common.ErrCannotGetEntity(foodratingmodel.EntityName, err)
	}

	if foodrating.Status == 0 {
		return common.ErrCannotGetEntity(foodratingmodel.EntityName, errors.New("foodrating not found"))
	}

	if err := biz.store.DeleteFoodRating(ctx, conditions); err != nil {
		return err
	}

	return nil
}
