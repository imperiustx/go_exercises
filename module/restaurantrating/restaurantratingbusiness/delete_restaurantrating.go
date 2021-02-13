package restaurantratingbusiness

import (
	"context"
	"errors"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/restaurantrating/restaurantratingmodel"
)

// DeleteRestaurantRatingStorage delete
type DeleteRestaurantRatingStorage interface {
	FindRestaurantRating(
		ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string) (*restaurantratingmodel.RestaurantRating, error)
	DeleteRestaurantRating(
		ctx context.Context,
		conditions map[string]interface{}) error
}

type deleteRestaurantRating struct {
	store DeleteRestaurantRatingStorage
}

// NewDeleteRestaurantRatingBiz delete
func NewDeleteRestaurantRatingBiz(store DeleteRestaurantRatingStorage) *deleteRestaurantRating {
	return &deleteRestaurantRating{store: store}
}

func (biz *deleteRestaurantRating) DeleteRestaurantRating(
	ctx context.Context,
	conditions map[string]interface{}) error {

	restaurantrating, err := biz.store.FindRestaurantRating(ctx, conditions)
	if err != nil {
		return common.ErrCannotGetEntity(restaurantratingmodel.EntityName, err)
	}

	if restaurantrating.Status == 0 {
		return common.ErrCannotGetEntity(
			restaurantratingmodel.EntityName, 
			errors.New("restaurantrating note found"),
		)
	}

	if err := biz.store.DeleteRestaurantRating(ctx, conditions); err != nil {
		return common.ErrCannotDeleteEntity(restaurantratingmodel.EntityName, err)
	}

	return nil
}
