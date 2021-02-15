package foodstorage

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/food/foodmodel"
)

func (s *sqlStore) ListFood(
	ctx context.Context,
	filter *foodmodel.Filter,
	paging *common.Paging,
	order *common.OrderSort,
	moreKeys ...string) ([]foodmodel.Food, error) {

	var foods []foodmodel.Food
	db := s.db.Table(foodmodel.Food{}.TableName())

	db = db.Where("status not in (0)")

	if f := filter; f != nil {
		if f.RestaurantID != 0 {
			db = db.Where("restaurant_id = ?", f.RestaurantID)
		}
		if f.CategoryID != 0 {
			db = db.Where("category_id = ?", f.CategoryID)
		}
		if f.Price != 0 {
			db = db.Where("price = ?", f.Price)
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

	if err := db.Find(&foods).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return foods, nil
}
