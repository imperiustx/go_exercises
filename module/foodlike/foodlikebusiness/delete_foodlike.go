package foodlikebusiness

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/foodlike/foodlikemodel"
)

// DeleteFoodLikeStorage delete
type DeleteFoodLikeStorage interface {
	FindFoodLike(ctx context.Context, uid, fid int) (*foodlikemodel.FoodLike, error)
	DeleteFoodLike(uid, fid int) error
}

type deleteFoodLike struct {
	store DeleteFoodLikeStorage
}

// NewDeleteFoodLikeBiz delete
func NewDeleteFoodLikeBiz(store DeleteFoodLikeStorage) *deleteFoodLike {
	return &deleteFoodLike{store: store}
}

func (biz *deleteFoodLike) DeleteFoodLike(ctx context.Context, uid, fid int) error {
	_, err := biz.store.FindFoodLike(ctx, uid, fid)
	if err != nil {
		return common.ErrCannotGetEntity(foodlikemodel.EntityName, err)
	}
	// TODO: consider add status to model
	// if foodlike.Status == 0 {
	// 	return common.ErrCannotGetEntity(foodlikemodel.EntityName, errors.New("foodlike not found"))
	// }

	if err := biz.store.DeleteFoodLike(uid, fid); err != nil {
		return err
	}

	return nil
}
