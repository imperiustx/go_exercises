package orderdetailstorage

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/orderdetail/orderdetailmodel"
)

func (s *sqlStore) ListOrderDetail(ctx context.Context, paging *common.Paging) ([]orderdetailmodel.OrderDetail, error) {
	db := s.db
	var orderdetails []orderdetailmodel.OrderDetail

	db = db.Table(orderdetailmodel.OrderDetail{}.TableName()).Where("status not in (0)")

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
	if err := db.Order("id asc").Find(&orderdetails).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return orderdetails, nil
}
