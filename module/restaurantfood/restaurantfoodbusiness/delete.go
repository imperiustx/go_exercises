package restaurantfoodbusiness

import (
	"context"
	"errors"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/restaurantfood/restaurantfoodmodel"
)

// DeleteRestaurantFoodStorage delete
type DeleteRestaurantFoodStorage interface {
	FindRestaurantFood(
		ctx context.Context,
		rid, fid int,
		moreInfo ...string) (*restaurantfoodmodel.RestaurantFood, error)
	DeleteRestaurantFood(
		ctx context.Context,
		rid, fid int) error
}

type deleteRestaurantFood struct {
	store DeleteRestaurantFoodStorage
}

// NewDeleteRestaurantFoodBiz delete
func NewDeleteRestaurantFoodBiz(store DeleteRestaurantFoodStorage) *deleteRestaurantFood {
	return &deleteRestaurantFood{store: store}
}

func (biz *deleteRestaurantFood) DeleteRestaurantFood(ctx context.Context, rid, fid int) error {
	restaurantfood, err := biz.store.FindRestaurantFood(ctx, rid, fid)
	if err != nil {
		return common.ErrCannotGetEntity(restaurantfoodmodel.EntityName, err)
	}

	if restaurantfood.Status == 0 {
		return common.ErrCannotGetEntity(
			restaurantfoodmodel.EntityName, 
			errors.New("category restaurant not found"),
		)
	}

	if err := biz.store.DeleteRestaurantFood(ctx, rid, fid); err != nil {
		return err
	}

	return nil
}
