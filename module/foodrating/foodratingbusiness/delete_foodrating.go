package foodratingbusiness

import (
	"context"
	"errors"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/foodrating/foodratingmodel"
)

// DeleteFoodRatingStorage delete
type DeleteFoodRatingStorage interface {
	FindFoodRating(ctx context.Context, id int) (*foodratingmodel.FoodRating, error)
	DeleteFoodRating(id int) error
}

type deleteFoodRating struct {
	store DeleteFoodRatingStorage
}

// NewDeleteFoodRatingBiz delete
func NewDeleteFoodRatingBiz(store DeleteFoodRatingStorage) *deleteFoodRating {
	return &deleteFoodRating{store: store}
}

func (biz *deleteFoodRating) DeleteFoodRating(ctx context.Context, id int) error {
	foodrating, err := biz.store.FindFoodRating(ctx, id)
	if err != nil {
		return common.ErrCannotGetEntity(foodratingmodel.EntityName, err)
	}

	if foodrating.Status == 0 {
		return common.ErrCannotGetEntity(foodratingmodel.EntityName, errors.New("foodrating not found"))
	}

	if err := biz.store.DeleteFoodRating(id); err != nil {
		return err
	}

	return nil
}
