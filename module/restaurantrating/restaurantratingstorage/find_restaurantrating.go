package restaurantratingstorage

import (
	"context"

	"github.com/imperiustx/go_excercises/module/restaurantrating/restaurantratingmodel"
)

func (s *sqlStore) FindRestaurantRating(ctx context.Context, id int) (*restaurantratingmodel.RestaurantRating, error) {
	db := s.db
	var restaurantrating restaurantratingmodel.RestaurantRating

	if err := db.Where("id = ?", id).
		First(&restaurantrating).Error; err != nil {
		return nil, err
	}

	return &restaurantrating, nil
}
