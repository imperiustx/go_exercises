package restaurantratingbusiness

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/restaurantrating/restaurantratingmodel"
)

// CreateRestaurantRatingStorage create
type CreateRestaurantRatingStorage interface {
	CreateRestaurantRating(ctx context.Context, data *restaurantratingmodel.RestaurantRatingCreate) error
}

type createRestaurantRating struct {
	store CreateRestaurantRatingStorage
}

// NewCreateRestaurantRatingBiz create
func NewCreateRestaurantRatingBiz(store CreateRestaurantRatingStorage) *createRestaurantRating {
	return &createRestaurantRating{store: store}
}

func (biz *createRestaurantRating) CreateNewRestaurantRating(ctx context.Context, data *restaurantratingmodel.RestaurantRatingCreate) error {
	if err := biz.store.CreateRestaurantRating(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(restaurantratingmodel.EntityName, err)
	}

	return nil
}
