package foodlikebusiness

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/foodlike/foodlikemodel"
)

// CreateFoodLikeStorage create
type CreateFoodLikeStorage interface {
	CreateFoodLike(ctx context.Context, data *foodlikemodel.FoodLikeCreate) error
}

type createFoodLike struct {
	store CreateFoodLikeStorage
}

// NewCreateFoodLikeBiz create
func NewCreateFoodLikeBiz(store CreateFoodLikeStorage) *createFoodLike {
	return &createFoodLike{store: store}
}

func (biz *createFoodLike) CreateNewFoodLike(ctx context.Context, data *foodlikemodel.FoodLikeCreate) error {
	if err := biz.store.CreateFoodLike(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(foodlikemodel.EntityName, err)
	}

	return nil
}
