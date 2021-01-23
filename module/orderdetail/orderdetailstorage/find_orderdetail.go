package orderdetailstorage

import (
	"context"

	"github.com/imperiustx/go_excercises/module/orderdetail/orderdetailmodel"
)

func (s *sqlStore) FindOrderDetail(ctx context.Context, id int) (*orderdetailmodel.OrderDetail, error) {
	db := s.db
	var orderdetail orderdetailmodel.OrderDetail

	if err := db.Where("id = ?", id).
		First(&orderdetail).Error; err != nil {
		return nil, err
	}

	return &orderdetail, nil
}
