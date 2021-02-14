package foodratingstorage

import (
	"context"

	"github.com/imperiustx/go_excercises/module/foodrating/foodratingmodel"
)

func (s *sqlStore) DeleteFoodRating(
	ctx context.Context,
	conditions map[string]interface{}) error {
		
	db := s.db.Table(foodratingmodel.FoodRating{}.TableName())

	if err := db.Where(conditions).Update("status", 0).Error; err != nil {
		return err
	}

	return nil
}
