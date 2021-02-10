package ordertrackingstorage

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/ordertracking/ordertrackingmodel"
)

func (s *sqlStore) CreateOrderTracking(ctx context.Context, data *ordertrackingmodel.OrderTrackingCreate) error {
	db := s.db.Table(data.TableName())

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
