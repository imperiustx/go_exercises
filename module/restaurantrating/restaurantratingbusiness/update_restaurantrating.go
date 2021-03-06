package restaurantratingbusiness

import (
	"context"
	"errors"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/restaurantrating/restaurantratingmodel"
)

// UpdateRestaurantRatingStorage update
type UpdateRestaurantRatingStorage interface {
	FindRestaurantRating(
		ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string) (*restaurantratingmodel.RestaurantRating, error)
	UpdateRestaurantRating(
		ctx context.Context, 
		conditions map[string]interface{},
		data *restaurantratingmodel.RestaurantRatingUpdate) error
}

type updateRestaurantRating struct {
	store UpdateRestaurantRatingStorage
}

// NewUpdateRestaurantRatingBiz update
func NewUpdateRestaurantRatingBiz(store UpdateRestaurantRatingStorage) *updateRestaurantRating {
	return &updateRestaurantRating{store: store}
}

func (biz *updateRestaurantRating) UpdateRestaurantRating(
	ctx context.Context, 
	conditions map[string]interface{},
	data *restaurantratingmodel.RestaurantRatingUpdate) error {

	restaurantrating, err := biz.store.FindRestaurantRating(ctx, conditions)
	if err != nil {
		return common.ErrCannotGetEntity(restaurantratingmodel.EntityName, err)
	}

	if restaurantrating.Status == 0 {
		return common.ErrCannotGetEntity(
			restaurantratingmodel.EntityName, 
			errors.New("restaurantrating not found"),
		)
	}

	if err := biz.store.UpdateRestaurantRating(ctx, conditions, data); err != nil {
		return common.ErrCannotUpdateEntity(restaurantratingmodel.EntityName, err)
	}

	return nil
}
