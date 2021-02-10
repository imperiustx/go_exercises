package ordertrackingstorage

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/ordertracking/ordertrackingmodel"
)

func (s *sqlStore) UpdateOrderTracking(ctx context.Context, id int, data *ordertrackingmodel.OrderTrackingUpdate) error {
	db := s.db.Table(data.TableName())

	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
