package restaurantratingbusiness

import (
	"context"
	"errors"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/restaurantrating/restaurantratingmodel"
)

// DeleteRestaurantRatingStorage delete
type DeleteRestaurantRatingStorage interface {
	FindRestaurantRating(ctx context.Context, id int) (*restaurantratingmodel.RestaurantRating, error)
	DeleteRestaurantRating(id int) error
}

type deleteRestaurantRating struct {
	store DeleteRestaurantRatingStorage
}

// NewDeleteRestaurantRatingBiz delete
func NewDeleteRestaurantRatingBiz(store DeleteRestaurantRatingStorage) *deleteRestaurantRating {
	return &deleteRestaurantRating{store: store}
}

func (biz *deleteRestaurantRating) DeleteRestaurantRating(ctx context.Context, id int) error {
	restaurantrating, err := biz.store.FindRestaurantRating(ctx, id)
	if err != nil {
		return common.ErrCannotGetEntity(restaurantratingmodel.EntityName, err)
	}

	if restaurantrating.Status == 0 {
		return common.ErrCannotGetEntity(restaurantratingmodel.EntityName, errors.New("restaurantrating note found"))
	}

	if err := biz.store.DeleteRestaurantRating(id); err != nil {
		return err
	}

	return nil
}
