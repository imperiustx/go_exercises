package foodlikebusiness

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/foodlike/foodlikemodel"
)

// GetFoodLikeStorage get
type GetFoodLikeStorage interface {
	FindFoodLike(
		ctx context.Context, 
		uid, fid int,
		moreInfo ...string) (*foodlikemodel.FoodLike, error)
}

type getFoodLikeBiz struct {
	store GetFoodLikeStorage
}

// NewGetFoodLikeBiz get
func NewGetFoodLikeBiz(store GetFoodLikeStorage) *getFoodLikeBiz {
	return &getFoodLikeBiz{store: store}
}

func (biz *getFoodLikeBiz) GetFoodLike(
	ctx context.Context, 
	uid, fid int,
	moreInfo ...string) (*foodlikemodel.FoodLike, error) {
		
	foodlike, err := biz.store.FindFoodLike(ctx, uid, fid)
	if err != nil {
		return nil, common.ErrCannotGetEntity(foodlikemodel.EntityName, err)
	}

	return foodlike, nil
}
