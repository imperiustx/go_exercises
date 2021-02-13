package orderstorage

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/order/ordermodel"
)

func (s *sqlStore) ListOrder(
	ctx context.Context, 
	paging *common.Paging,
	order *common.OrderSort,
	moreKeys ...string) ([]ordermodel.Order, error) {
		
	db := s.db.Table(ordermodel.Order{}.TableName())
	var orders []ordermodel.Order

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

	if err := db.Find(&orders).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return orders, nil
}
