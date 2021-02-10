package ordertrackingstorage

import "github.com/imperiustx/go_excercises/module/ordertracking/ordertrackingmodel"

func (s *sqlStore) DeleteOrderTracking(id int) error {
	db := s.db.Table(ordertrackingmodel.OrderTracking{}.TableName())

	if err := db.Where("id = ?", id).
		Update("status", 0).Error; err != nil {
		return err
	}

	return nil
}
