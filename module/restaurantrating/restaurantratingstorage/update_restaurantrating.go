package restaurantratingstorage

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/restaurantrating/restaurantratingmodel"
)

func (s *sqlStore) UpdateRestaurantRating(ctx context.Context, id int, data *restaurantratingmodel.RestaurantRatingUpdate) error {
	db := s.db

	if err := db.Table(data.TableName()).
		Where("id = ?", id).
		Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
