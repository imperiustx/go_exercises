package useraddressstorage

func (s *sqlStore) UpdateUserAddress(id string, v interface{}) error {
	db := s.db

	if err := db.Where("id = ?", id).Updates(v).Error; err != nil {
		return err
	}

	return nil
}
