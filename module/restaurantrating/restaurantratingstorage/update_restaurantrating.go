package restaurantratingstorage

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/restaurantrating/restaurantratingmodel"
)

func (s *sqlStore) UpdateRestaurantRating(
	ctx context.Context,
	conditions map[string]interface{},
	data *restaurantratingmodel.RestaurantRatingUpdate) error {

	db := s.db.Table(data.TableName())

	if err := db.Where(conditions).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
