package orderdetailstorage

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/orderdetail/orderdetailmodel"
)

func (s *sqlStore) UpdateOrderDetail(
	ctx context.Context, 
	conditions map[string]interface{},
	data *orderdetailmodel.OrderDetailUpdate) error {

	db := s.db.Table(data.TableName())

	if err := db.Where(conditions).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
