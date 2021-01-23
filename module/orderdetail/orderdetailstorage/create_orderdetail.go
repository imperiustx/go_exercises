package orderdetailstorage

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/orderdetail/orderdetailmodel"
)

func (s *sqlStore) CreateOrderDetail(ctx context.Context, data *orderdetailmodel.OrderDetailCreate) error {
	db := s.db

	if err := db.Table(data.TableName()).
		Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
