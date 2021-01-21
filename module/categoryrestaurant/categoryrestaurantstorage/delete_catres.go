package categoryrestaurantstorage

import "github.com/imperiustx/go_excercises/module/categoryrestaurant/categoryrestaurantmodel"

func (s *sqlStore) DeleteCategoryRestaurant(cid, rid int) error {
	db := s.db

	if err := db.Table("category_restaurants").
		Where(&categoryrestaurantmodel.CategoryRestaurant{
			CategoryID:   cid,
			RestaurantID: rid,
		}).Update("status", 0).Error; err != nil {
		return err
	}

	return nil
}
