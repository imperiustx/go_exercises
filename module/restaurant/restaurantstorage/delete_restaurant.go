package restaurantstorage

func (s *sqlStore) DeleteRestaurant(id string) error {
	db := s.db

	if err := db.Delete(id).Error; err != nil {
		return err
	}

	return nil
}
