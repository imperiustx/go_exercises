package orderdetailstorage

import (
	"context"

	"github.com/imperiustx/go_excercises/module/orderdetail/orderdetailmodel"
)

func (s *sqlStore) DeleteOrderDetail(ctx context.Context, conditions map[string]interface{}) error {
	db := s.db.Table(orderdetailmodel.OrderDetail{}.TableName())

	if err := db.Where(conditions).Update("status", 0).Error; err != nil {
		return err
	}

	return nil
}
