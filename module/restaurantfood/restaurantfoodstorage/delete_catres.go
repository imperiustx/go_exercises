package restaurantfoodstorage

import (
	"context"

	"github.com/imperiustx/go_excercises/module/restaurantfood/restaurantfoodmodel"
)

func (s *sqlStore) DeleteRestaurantFood(ctx context.Context, rid, fid int) error {
	db := s.db.Table(restaurantfoodmodel.RestaurantFood{}.TableName())

	if err := db.Where(&restaurantfoodmodel.RestaurantFood{
		RestaurantID: rid,
		FoodID:       fid,
	}).Update("status", 0).Error; err != nil {
		return err
	}

	return nil
}
