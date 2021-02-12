package restaurantstorage

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/restaurant/restaurantmodel"
)

func (s *sqlStore) ListRestaurant(
	ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
	order *common.OrderSort,
	moreKeys ...string) ([]restaurantmodel.Restaurant, error) {
	db := s.db.Table(restaurantmodel.Restaurant{}.TableName())
	var restaurants []restaurantmodel.Restaurant

	db = db.Where("status not in (0)")

	if f := filter; f != nil {
		if f.OwnerID != "" {
			db = db.Where("owner_id= ?", f.OwnerID)
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

	if err := db.Find(&restaurants).Error; err != nil {
		return nil, err
	}

	return restaurants, nil
}
