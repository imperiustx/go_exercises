package orderdetailstorage

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/orderdetail/orderdetailmodel"
)

func (s *sqlStore) ListOrderDetail(
	ctx context.Context, 
	filter *orderdetailmodel.Filter,
	paging *common.Paging,
	order *common.OrderSort,
	moreKeys ...string) ([]orderdetailmodel.OrderDetail, error) {

	db := s.db.Table(orderdetailmodel.OrderDetail{}.TableName())
	var orderdetails []orderdetailmodel.OrderDetail

	db = db.Where("status not in (0)")

	if f := filter; f != nil {
		if f.Price != 0 {
			db = db.Where("price = ?", f.Price)
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

	if err := db.Find(&orderdetails).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return orderdetails, nil
}
