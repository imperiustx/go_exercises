package foodstorage

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/food/foodmodel"
)

func (s *sqlStore) ListFood(ctx context.Context, paging *common.Paging) ([]foodmodel.Food, error) {
	db := s.db
	var foods []foodmodel.Food

	db = db.Table(foodmodel.Food{}.TableName()).Where("status not in (0)")

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
	if err := db.Order("id asc").Find(&foods).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return foods, nil
}
