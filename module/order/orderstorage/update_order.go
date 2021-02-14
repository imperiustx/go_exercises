package orderstorage

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/order/ordermodel"
)

func (s *sqlStore) UpdateOrder(
	ctx context.Context,
	conditions map[string]interface{},
	data *ordermodel.OrderUpdate) error {

	db := s.db.Table(data.TableName())

	if err := db.Where(conditions).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
