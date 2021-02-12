package restaurantstorage

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/restaurant/restaurantmodel"
)

func (s *sqlStore) UpdateRestaurant(
	ctx context.Context, 
	conditions map[string]interface{}, 
	data *restaurantmodel.RestaurantUpdate) error {
	db := s.db.Table(data.TableName())

	if err := db.Where(conditions).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
