package orderdetailstorage

func (s *sqlStore) DeleteOrderDetail(id int) error {
	db := s.db

	if err := db.Table("orderdetails").
		Where("id = ?", id).
		Update("status", 0).Error; err != nil {
		return err
	}

	return nil
}
