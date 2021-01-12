package userstorage

func (s *sqlStore) CreateUser(v interface{}) error {
	db := s.db

	if err := db.Create(v).Error; err != nil {
		return err
	}

	return nil
}
