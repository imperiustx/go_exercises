package restaurantstorage

import "github.com/imperiustx/go_excercises/module/restaurant/restaurantmodel"

func (s *sqlStore) DeleteRestaurant(conditions map[string]interface{}) error {
	db := s.db.Table(restaurantmodel.Restaurant{}.TableName())

	if err := db.Where("id = ?", conditions["id"]).
		Update("status", 0).Error; err != nil {
		return err
	}

	return nil
}
