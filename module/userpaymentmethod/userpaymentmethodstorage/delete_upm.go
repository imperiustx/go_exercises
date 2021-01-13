package userpaymentmethodstorage

func (s *sqlStore) DeleteUserPaymentMethod(id string) error {
	db := s.db

	if err := db.Delete("userpaymentmethods", id).Error; err != nil {
		return err
	}

	return nil
}
