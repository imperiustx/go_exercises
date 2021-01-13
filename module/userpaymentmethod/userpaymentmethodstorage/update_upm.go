package userpaymentmethodstorage

func (s *sqlStore) UpdateUserPaymentMethod(id string, v interface{}) error {
	db := s.db

	if err := db.Where("id = ?", id).Updates(v).Error; err != nil {
		return err
	}

	return nil
}
