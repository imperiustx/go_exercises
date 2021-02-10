package orderstorage

import (
	"context"

	"github.com/imperiustx/go_excercises/module/order/ordermodel"
)

func (s *sqlStore) FindOrder(ctx context.Context, id int, moreInfo ...string) (*ordermodel.Order, error) {
	db := s.db.Table(ordermodel.Order{}.TableName())

	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}

	var order ordermodel.Order

	if err := db.Where("id = ?", id).First(&order).Error; err != nil {
		return nil, err
	}

	return &order, nil
}
