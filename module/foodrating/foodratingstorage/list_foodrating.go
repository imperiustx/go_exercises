package foodratingstorage

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/foodrating/foodratingmodel"
)

func (s *sqlStore) ListFoodRating(
	ctx context.Context,
	filter *foodratingmodel.Filter,
	paging *common.Paging,
	order *common.OrderSort,
	moreKeys ...string) ([]foodratingmodel.FoodRating, error) {

	db := s.db.Table(foodratingmodel.FoodRating{}.TableName())
	var foodratings []foodratingmodel.FoodRating

	db = db.Where("status not in (0)")

	if f := filter; f != nil {
		if f.Point != 0 {
			db = db.Where("point = ?", f.Point)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	db = db.Limit(paging.Limit)

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

	if err := db.Find(&foodratings).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return foodratings, nil
}
