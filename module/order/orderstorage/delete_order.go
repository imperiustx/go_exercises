package orderstorage

import "github.com/imperiustx/go_excercises/module/order/ordermodel"

func (s *sqlStore) DeleteOrder(id int) error {
	db := s.db.Table(ordermodel.Order{}.TableName())

	if err := db.Where("id = ?", id).
		Update("status", 0).Error; err != nil {
		return err
	}

	return nil
}
