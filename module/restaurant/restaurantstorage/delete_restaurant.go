package restaurantstorage

func (s *sqlStore) DeleteRestaurant(id int) error {
	db := s.db

	if err := db.Table("restaurants").
		Where("id = ?", id).
		Update("status", 0).Error; err != nil {
		return err
	}

	return nil
}
