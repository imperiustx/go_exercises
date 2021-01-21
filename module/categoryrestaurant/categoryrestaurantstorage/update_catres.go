package categoryrestaurantstorage

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/categoryrestaurant/categoryrestaurantmodel"
)

func (s *sqlStore) UpdateCategoryRestaurant(ctx context.Context, cid, rid int, data *categoryrestaurantmodel.CategoryRestaurantUpdate) error {
	db := s.db

	if err := db.Table(data.TableName()).
		Where(&categoryrestaurantmodel.CategoryRestaurant{
			CategoryID:   cid,
			RestaurantID: rid,
		}).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
