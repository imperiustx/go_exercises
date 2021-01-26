package userstorage

func (s *sqlStore) ReactiveUser(conditions map[string]interface{}) error {
	db := s.db

	if err := db.Table("users").
		Where("id = ?", conditions["id"]).
		Update("status", 1).Error; err != nil {
		return err
	}

	return nil
}
