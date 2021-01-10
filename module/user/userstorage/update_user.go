package userstorage

func (s *sqlStore) UpdateUser(id string, v interface{}) error {
	db := s.db

	if err := db.Where("id = ?", id).Updates(v).Error; err != nil {
		return err
	}

	return nil
}
