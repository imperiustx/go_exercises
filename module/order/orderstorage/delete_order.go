package orderstorage

func (s *sqlStore) DeleteOrder(id int) error {
	db := s.db

	if err := db.Table("orders").
		Where("id = ?", id).
		Update("status", 0).Error; err != nil {
		return err
	}

	return nil
}
