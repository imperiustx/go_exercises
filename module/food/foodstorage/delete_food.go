package foodstorage

func (s *sqlStore) DeleteFood(id int) error {
	db := s.db

	if err := db.Table("foods").
		Where("id = ?", id).
		Update("status", 0).Error; err != nil {
		return err
	}

	return nil
}
