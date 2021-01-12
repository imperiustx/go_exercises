package restaurantstorage

import "github.com/imperiustx/go_excercises/module/restaurant/restaurantmodel"

func (s *sqlStore) ListRestaurant() ([]restaurantmodel.Restaurant, error) {
	db := s.db
	var restaurants []restaurantmodel.Restaurant

	if err := db.Find(&restaurants).Error; err != nil {
		return nil, err
	}

	return restaurants, nil
}
