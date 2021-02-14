package foodratingbusiness

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/foodrating/foodratingmodel"
)

// CreateFoodRatingStorage create
type CreateFoodRatingStorage interface {
	FindFoodRating(
		ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string) (*foodratingmodel.FoodRating, error)
	CreateFoodRating(
		ctx context.Context,
		data *foodratingmodel.FoodRatingCreate) error
}

type createFoodRating struct {
	store CreateFoodRatingStorage
}

// NewCreateFoodRatingBiz create
func NewCreateFoodRatingBiz(store CreateFoodRatingStorage) *createFoodRating {
	return &createFoodRating{store: store}
}

func (biz *createFoodRating) CreateNewFoodRating(
	ctx context.Context,
	data *foodratingmodel.FoodRatingCreate) error {

	rating, err := biz.store.FindFoodRating(ctx, map[string]interface{}{"id": data.ID})
	if rating != nil {
		return common.ErrEntityExisted(foodratingmodel.EntityName, err)
	}
	
	if err := biz.store.CreateFoodRating(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(foodratingmodel.EntityName, err)
	}

	return nil
}
