package citystorage

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/city/citymodel"
)

func (s *sqlStore) ListCity(ctx context.Context, paging *common.Paging) ([]citymodel.City, error) {
	db := s.db.Table(citymodel.City{}.TableName())
	var cities []citymodel.City

	db = db.Where("status not in (0)")

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
	if err := db.Order("id asc").Find(&cities).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return cities, nil
}
