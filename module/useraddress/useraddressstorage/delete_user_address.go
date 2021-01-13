package useraddressstorage

func (s *sqlStore) DeleteUserAddress(id string) error {
	db := s.db

	if err := db.Delete("useraddresss", id).Error; err != nil {
		return err
	}

	return nil
}
