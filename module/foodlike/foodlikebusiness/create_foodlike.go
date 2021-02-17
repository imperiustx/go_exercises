package foodlikebusiness

import (
	"context"
	"strconv"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/foodlike/foodlikemodel"
)

// CreateFoodLikeStorage create
type CreateFoodLikeStorage interface {
	FindFoodLike(
		ctx context.Context,
		uid, fid int,
		moreInfo ...string) (*foodlikemodel.FoodLike, error)
	CreateFoodLike(
		ctx context.Context,
		data *foodlikemodel.FoodLikeCreate) error
}

type createFoodLike struct {
	store CreateFoodLikeStorage
}

// NewCreateFoodLikeBiz create
func NewCreateFoodLikeBiz(store CreateFoodLikeStorage) *createFoodLike {
	return &createFoodLike{store: store}
}

func (biz *createFoodLike) CreateNewFoodLike(
	ctx context.Context,
	data *foodlikemodel.FoodLikeCreate) error {

	uid, err := strconv.Atoi(data.UserID)
	if err != nil {
		return common.ErrCannotConvertGivenData(data.UserID, err)
	}
	fid, err := strconv.Atoi(data.FoodID)
	if err != nil {
		return common.ErrCannotConvertGivenData(data.FoodID, err)
	}

	like, err := biz.store.FindFoodLike(ctx, uid, fid)
	if like != nil {
		return common.ErrEntityExisted(foodlikemodel.EntityName, err)
	}

	if err := biz.store.CreateFoodLike(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(foodlikemodel.EntityName, err)
	}

	return nil
}
