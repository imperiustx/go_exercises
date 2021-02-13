package restaurantratingstorage

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/restaurantrating/restaurantratingmodel"
)

func (s *sqlStore) ListRestaurantRating(
	ctx context.Context,
	filter *restaurantratingmodel.Filter,
	paging *common.Paging,
	order *common.OrderSort,
	moreKeys ...string) ([]restaurantratingmodel.RestaurantRating, error) {

	db := s.db.Table(restaurantratingmodel.RestaurantRating{}.TableName())
	var restaurantratings []restaurantratingmodel.RestaurantRating

	db = db.Where("status not in (0)")

	if f := filter; f != nil {
		if f.Point != 0 {
			db = db.Where("point = ?", f.Point)
		}
		if f.UserID != 0 {
			db = db.Where("user_id = ?", f.UserID)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	db = db.Limit(paging.Limit)

	for _, k := range moreKeys {
		db = db.Preload(k)
	}

	if paging.Cursor > 0 {
		db = db.Where("id < ?", paging.Cursor)
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	if o := order; o != nil {
		if o.Order == "asc" {
			db = db.Order("id asc")
		}
		if o.Order == "desc" {
			db = db.Order("id desc")
		}
	}

	if err := db.Find(&restaurantratings).Error; err != nil {
		return nil, err
	}

	return restaurantratings, nil
}
