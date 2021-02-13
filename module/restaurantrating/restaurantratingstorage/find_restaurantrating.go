package restaurantratingstorage

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/restaurantrating/restaurantratingmodel"
)

func (s *sqlStore) FindRestaurantRating(
	ctx context.Context,
	conditions map[string]interface{},
	moreInfo ...string) (*restaurantratingmodel.RestaurantRating, error) {

	db := s.db.Table(restaurantratingmodel.RestaurantRating{}.TableName())

	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}

	var restaurantrating restaurantratingmodel.RestaurantRating

	if err := db.Where(conditions).First(&restaurantrating).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return &restaurantrating, nil
}
