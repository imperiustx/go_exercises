package restaurantstorage

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/restaurant/restaurantmodel"
)

func (s *sqlStore) ListRestaurant(ctx context.Context, paging *common.Paging) ([]restaurantmodel.Restaurant, error) {
	db := s.db
	var restaurants []restaurantmodel.Restaurant

	db = db.Table(restaurantmodel.Restaurant{}.TableName()).Where("status not in (0)")

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	db = db.Limit(paging.Limit)

	if paging.Cursor > 0 {
		db = db.Where("id < ?", paging.Cursor)
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	if err := db.Order("id desc").Find(&restaurants).Error; err != nil {
		return nil, err
	}

	return restaurants, nil
}
