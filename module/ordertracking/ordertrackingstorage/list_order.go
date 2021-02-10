package ordertrackingstorage

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/ordertracking/ordertrackingmodel"
)

func (s *sqlStore) ListOrderTracking(
	ctx context.Context,
	paging *common.Paging,
	ordertracking *common.OrderSort,
	moreKeys ...string) ([]ordertrackingmodel.OrderTracking, error) {
	db := s.db.Table(ordertrackingmodel.OrderTracking{}.TableName())
	var ordertrackings []ordertrackingmodel.OrderTracking

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

	if o := ordertracking; o != nil {
		if o.Order == "asc" {
			db = db.Order("id asc")
		}
		if o.Order == "desc" {
			db = db.Order("id desc")
		}
	}

	if err := db.Find(&ordertrackings).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return ordertrackings, nil
}
