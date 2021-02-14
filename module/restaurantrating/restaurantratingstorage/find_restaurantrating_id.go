package restaurantratingstorage

import (
	"context"

	"github.com/imperiustx/go_excercises/module/restaurantrating/restaurantratingmodel"
)

func (s *sqlStore) FindRestaurantRatingID(
	ctx context.Context,
	uid, rid int,
	moreInfo ...string) (*int, error) {

	db := s.db.Table(restaurantratingmodel.RestaurantRating{}.TableName())
	var restaurantrating restaurantratingmodel.RestaurantRating

	if err := db.Where(&restaurantratingmodel.RestaurantRating{
		UserID: uid,
		RestaurantID: rid,
	}).First(&restaurantrating).Error; err != nil {
		return nil, err
	}

	return &restaurantrating.ID, nil
}
