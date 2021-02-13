package restaurantratingstorage

import (
	"context"

	"github.com/imperiustx/go_excercises/module/restaurantrating/restaurantratingmodel"
)

func (s *sqlStore) DeleteRestaurantRating(
	ctx context.Context,
	conditions map[string]interface{}) error {

	db := s.db.Table(restaurantratingmodel.RestaurantRating{}.TableName())

	if err := db.Where(conditions).Update("status", 0).Error; err != nil {
		return err
	}

	return nil
}
