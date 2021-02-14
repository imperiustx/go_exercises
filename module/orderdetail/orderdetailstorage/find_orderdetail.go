package orderdetailstorage

import (
	"context"

	"github.com/imperiustx/go_excercises/module/orderdetail/orderdetailmodel"
)

func (s *sqlStore) FindOrderDetail(
	ctx context.Context, 
	conditions map[string]interface{}, 
	moreInfo ...string) (*orderdetailmodel.OrderDetail, error) {

	db := s.db.Table(orderdetailmodel.OrderDetail{}.TableName())
	var orderdetail orderdetailmodel.OrderDetail

	if err := db.Where(conditions).First(&orderdetail).Error; err != nil {
		return nil, err
	}

	return &orderdetail, nil
}
