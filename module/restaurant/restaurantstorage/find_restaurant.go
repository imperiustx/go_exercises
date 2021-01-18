package restaurantstorage

import (
	"context"

	"github.com/imperiustx/go_excercises/module/restaurant/restaurantmodel"
)

func (s *sqlStore) FindRestaurant(ctx context.Context, id int) (*restaurantmodel.Restaurant, error) {
	db := s.db
	var restaurant restaurantmodel.Restaurant

	if err := db.Where("id = ?", id).
		First(&restaurant).Error; err != nil {
		return nil, err
	}

	return &restaurant, nil
}
