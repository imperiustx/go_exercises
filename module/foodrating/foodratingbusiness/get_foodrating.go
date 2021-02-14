package foodratingbusiness

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/foodrating/foodratingmodel"
)

// GetFoodRatingStorage get
type GetFoodRatingStorage interface {
	FindFoodRating(
		ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string) (*foodratingmodel.FoodRating, error)
}

type getFoodRatingBiz struct {
	store GetFoodRatingStorage
}

// NewGetFoodRatingBiz get
func NewGetFoodRatingBiz(store GetFoodRatingStorage) *getFoodRatingBiz {
	return &getFoodRatingBiz{store: store}
}

func (biz *getFoodRatingBiz) GetFoodRating(
	ctx context.Context,
	conditions map[string]interface{},
	moreInfo ...string) (*foodratingmodel.FoodRating, error) {

	foodrating, err := biz.store.FindFoodRating(ctx, conditions)
	if err != nil {
		return nil, common.ErrCannotGetEntity(foodratingmodel.EntityName, err)
	}

	return foodrating, nil
}
