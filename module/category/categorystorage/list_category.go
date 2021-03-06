package categorystorage

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/category/categorymodel"
)

func (s *sqlStore) ListCategory(
	ctx context.Context,
	paging *common.Paging,
	order *common.OrderSort,
	moreKeys ...string) ([]categorymodel.Category, error) {

	db := s.db.Table(categorymodel.Category{}.TableName())
	var categories []categorymodel.Category

	db = db.Where("status not in (0)")

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

	if err := db.Find(&categories).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return categories, nil
}
