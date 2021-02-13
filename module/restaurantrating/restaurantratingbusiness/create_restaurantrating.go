package restaurantratingbusiness

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/restaurantrating/restaurantratingmodel"
)

// CreateRestaurantRatingStorage create
type CreateRestaurantRatingStorage interface {
	FindRestaurantRating(
		ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string) (*restaurantratingmodel.RestaurantRating, error)
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

	// TODO: think again the condition below
	rating, err := biz.store.FindRestaurantRating(ctx, map[string]interface{}{"id": data.ID})
	if rating != nil {
		return common.ErrEntityExisted(restaurantratingmodel.EntityName, err)
	}

	if err := biz.store.CreateRestaurantRating(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(restaurantratingmodel.EntityName, err)
	}

	return nil
}
