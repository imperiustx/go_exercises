package orderdetailstorage

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/orderdetail/orderdetailmodel"
)

func (s *sqlStore) UpdateOrderDetail(ctx context.Context, id int, data *orderdetailmodel.OrderDetailUpdate) error {
	db := s.db

	if err := db.Table(data.TableName()).
		Where("id = ?", id).
		Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
