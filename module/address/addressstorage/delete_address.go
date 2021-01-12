package addressstorage

func (s *sqlStore) DeleteAddress(id string) error {
	db := s.db

	if err := db.Delete(id).Error; err != nil {
		return err
	}

	return nil
}
