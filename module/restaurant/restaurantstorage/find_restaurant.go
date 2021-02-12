package restaurantstorage

import (
	"context"

	"github.com/imperiustx/go_excercises/module/restaurant/restaurantmodel"
)

func (s *sqlStore) FindRestaurant(
	ctx context.Context,
	conditions map[string]interface{},
	moreInfo ...string) (*restaurantmodel.Restaurant, error) {

	db := s.db.Table(restaurantmodel.Restaurant{}.TableName())

	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}

	var restaurant restaurantmodel.Restaurant

	if err := db.Where(conditions).First(&restaurant).Error; err != nil {
		return nil, err
	}

	return &restaurant, nil
}
