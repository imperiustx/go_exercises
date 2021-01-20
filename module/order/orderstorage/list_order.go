package orderstorage

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/order/ordermodel"
)

func (s *sqlStore) ListOrder(ctx context.Context, paging *common.Paging) ([]ordermodel.Order, error) {
	db := s.db
	var orders []ordermodel.Order

	db = db.Table(ordermodel.Order{}.TableName()).Where("status not in (0)")

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
	if err := db.Order("id asc").Find(&orders).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return orders, nil
}
