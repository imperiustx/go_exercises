package categoryrestaurantstorage

import (
	"context"

	"github.com/imperiustx/go_excercises/module/categoryrestaurant/categoryrestaurantmodel"
)

func (s *sqlStore) FindCategoryRestaurant(ctx context.Context, rid, cid int) (*categoryrestaurantmodel.CategoryRestaurant, error) {
	db := s.db
	var categoryrestaurant categoryrestaurantmodel.CategoryRestaurant

	if err := db.Where(&categoryrestaurantmodel.CategoryRestaurant{
		CategoryID:   cid,
		RestaurantID: rid,
	}).First(&categoryrestaurant).Error; err != nil {
		return nil, err
	}

	return &categoryrestaurant, nil
}
