package foodratingstorage

import (
	"context"

	"github.com/imperiustx/go_excercises/module/foodrating/foodratingmodel"
)

func (s *sqlStore) FindFoodRating(ctx context.Context, id int) (*foodratingmodel.FoodRating, error) {
	db := s.db
	var foodrating foodratingmodel.FoodRating

	if err := db.Where("id = ?", id).
		First(&foodrating).Error; err != nil {
		return nil, err
	}

	return &foodrating, nil
}
