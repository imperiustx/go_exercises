package restaurantratingstorage

func (s *sqlStore) DeleteRestaurantRating(id int) error {
	db := s.db

	if err := db.Table("restaurant_ratings").
		Where("id = ?", id).
		Update("status", 0).Error; err != nil {
		return err
	}

	return nil
}
