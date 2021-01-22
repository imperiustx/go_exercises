package foodratingstorage

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/foodrating/foodratingmodel"
)

func (s *sqlStore) ListFoodRating(ctx context.Context, paging *common.Paging) ([]foodratingmodel.FoodRating, error) {
	db := s.db
	var foodratings []foodratingmodel.FoodRating

	db = db.Table(foodratingmodel.FoodRating{}.TableName()).Where("status not in (0)")

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	db = db.Limit(paging.Limit)

	if paging.Cursor > 0 {
		db = db.Where("id < ?", paging.Cursor)
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	// id desc
	if err := db.Order("id asc").Find(&foodratings).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return foodratings, nil
}
