package ordertrackingstorage

import (
	"context"

	"github.com/imperiustx/go_excercises/module/ordertracking/ordertrackingmodel"
)

func (s *sqlStore) DeleteOrderTracking(
	ctx context.Context,
	conditions map[string]interface{}) error {

	db := s.db.Table(ordertrackingmodel.OrderTracking{}.TableName())

	if err := db.Where(conditions).Update("status", 0).Error; err != nil {
		return err
	}

	return nil
}
