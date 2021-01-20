package categorystorage

func (s *sqlStore) DeleteCategory(id int) error {
	db := s.db

	if err := db.Table("categories").
		Where("id = ?", id).
		Update("status", 0).Error; err != nil {
		return err
	}

	return nil
}
