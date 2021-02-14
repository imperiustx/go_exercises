package orderstorage

import (
	"context"

	"github.com/imperiustx/go_excercises/module/order/ordermodel"
)

func (s *sqlStore) DeleteOrder(
	ctx context.Context,
	conditions map[string]interface{}) error {

	db := s.db.Table(ordermodel.Order{}.TableName())

	if err := db.Where(conditions).Update("status", 0).Error; err != nil {
		return err
	}

	return nil
}
