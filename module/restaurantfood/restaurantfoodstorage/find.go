package restaurantfoodstorage

import (
	"context"

	"github.com/imperiustx/go_excercises/module/restaurantfood/restaurantfoodmodel"
)

func (s *sqlStore) FindRestaurantFood(
	ctx context.Context,
	rid, fid int,
	moreInfo ...string) (*restaurantfoodmodel.RestaurantFood, error) {
	db := s.db.Table(restaurantfoodmodel.RestaurantFood{}.TableName())

	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}

	var restaurantfood restaurantfoodmodel.RestaurantFood

	if err := db.Where(&restaurantfoodmodel.RestaurantFood{
		RestaurantID: rid,
		FoodID:   fid,
	}).First(&restaurantfood).Error; err != nil {
		return nil, err
	}

	return &restaurantfood, nil
}
