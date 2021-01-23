package restaurantratingstorage

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/restaurantrating/restaurantratingmodel"
)

func (s *sqlStore) ListRestaurantRating(ctx context.Context, paging *common.Paging) ([]restaurantratingmodel.RestaurantRating, error) {
	db := s.db
	var restaurantratings []restaurantratingmodel.RestaurantRating

	db = db.Table(restaurantratingmodel.RestaurantRating{}.TableName()).Where("status not in (0)")

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	db = db.Limit(paging.Limit)

	if paging.Cursor > 0 {
		db = db.Where("id < ?", paging.Cursor)
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	if err := db.Order("id desc").Find(&restaurantratings).Error; err != nil {
		return nil, err
	}

	return restaurantratings, nil
}
