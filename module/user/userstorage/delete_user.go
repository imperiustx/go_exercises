package userstorage

func (s *sqlStore) DeleteUser(id string) error {
	db := s.db

	/*
	TODO: fix into this
	if err := db.Table("notes").
		Where("id = ?", id).
		Update("status", 0).Error; err != nil {
		return err
	*/
	if err := db.Delete("users", id).Error; err != nil {
		return err
	}

	return nil
}
