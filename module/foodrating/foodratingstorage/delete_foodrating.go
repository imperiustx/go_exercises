package foodratingstorage

func (s *sqlStore) DeleteFoodRating(id int) error {
	db := s.db

	if err := db.Table("food_ratings").
		Where("id = ?", id).
		Update("status", 0).Error; err != nil {
		return err
	}

	return nil
}
