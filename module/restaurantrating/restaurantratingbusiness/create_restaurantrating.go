package restaurantratingbusiness

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/restaurantrating/restaurantratingmodel"
)

// CreateRestaurantRatingStorage create
type CreateRestaurantRatingStorage interface {
	FindRestaurantRatingID(
		ctx context.Context,
		uid, rid int,
		moreInfo ...string) (*int, error)
	CreateRestaurantRating(
		ctx context.Context,
		data *restaurantratingmodel.RestaurantRatingCreate) error
}

type createRestaurantRating struct {
	store CreateRestaurantRatingStorage
}

// NewCreateRestaurantRatingBiz create
func NewCreateRestaurantRatingBiz(store CreateRestaurantRatingStorage) *createRestaurantRating {
	return &createRestaurantRating{store: store}
}

func (biz *createRestaurantRating) CreateNewRestaurantRating(
	ctx context.Context,
	data *restaurantratingmodel.RestaurantRatingCreate) error {

	rating, err := biz.store.FindRestaurantRatingID(ctx, data.UserID, data.RestaurantID)
	if rating != nil {
		return common.ErrEntityExisted(restaurantratingmodel.EntityName, err)
	}

	if err := biz.store.CreateRestaurantRating(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(restaurantratingmodel.EntityName, err)
	}

	return nil
}
