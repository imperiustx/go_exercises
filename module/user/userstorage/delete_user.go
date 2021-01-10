package userstorage

func (s *sqlStore) DeleteUser(id string) error {
	db := s.db

	if err := db.Delete("users", id).Error; err != nil {
		return err
	}

	return nil
}
