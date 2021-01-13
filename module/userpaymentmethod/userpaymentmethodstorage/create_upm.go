package userpaymentmethodstorage

func (s *sqlStore) CreateUserPaymentMethod(v interface{}) error {
	db := s.db

	if err := db.Create(v).Error; err != nil {
		return err
	}

	return nil
}
