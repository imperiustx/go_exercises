package foodratingstorage

import (
	"context"

	"github.com/imperiustx/go_excercises/module/foodrating/foodratingmodel"
)

func (s *sqlStore) FindFoodRating(
	ctx context.Context,
	conditions map[string]interface{},
	moreInfo ...string) (*foodratingmodel.FoodRating, error) {

	db := s.db.Table(foodratingmodel.FoodRating{}.TableName())
	var foodrating foodratingmodel.FoodRating

	if err := db.Where(conditions).First(&foodrating).Error; err != nil {
		return nil, err
	}

	return &foodrating, nil
}
