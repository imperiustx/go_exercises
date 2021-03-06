package restaurantratingbusiness

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/restaurantrating/restaurantratingmodel"
)

// GetRestaurantRatingStorage get
type GetRestaurantRatingStorage interface {
	FindRestaurantRating(
		ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string) (*restaurantratingmodel.RestaurantRating, error)
}

type getRestaurantRating struct {
	store GetRestaurantRatingStorage
}

// NewGetRestaurantRatingBiz get
func NewGetRestaurantRatingBiz(store GetRestaurantRatingStorage) *getRestaurantRating {
	return &getRestaurantRating{store: store}
}

func (biz *getRestaurantRating) GetRestaurantRating(
	ctx context.Context,
	conditions map[string]interface{},
	moreInfo ...string) (*restaurantratingmodel.RestaurantRating, error) {
		
	restaurantrating, err := biz.store.FindRestaurantRating(ctx, conditions)
	if err != nil {
		return nil, common.ErrCannotGetEntity(restaurantratingmodel.EntityName, err)
	}

	return restaurantrating, nil
}
