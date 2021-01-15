package userstorage

func (s *sqlStore) DeleteUser(id int) error {
	db := s.db

	if err := db.Table("users").
		Where("id = ?", id).
		Update("status", 0).Error; err != nil {
		return err
	}

	return nil
}
