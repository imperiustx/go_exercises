package paymentmethodstorage

func (s *sqlStore) DeletePaymentMethod(id string) error {
	db := s.db

	if err := db.Delete("paymentMethods", id).Error; err != nil {
		return err
	}

	return nil
}
