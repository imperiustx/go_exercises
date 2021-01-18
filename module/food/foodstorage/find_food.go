package foodstorage

import (
	"context"

	"github.com/imperiustx/go_excercises/module/food/foodmodel"
)

func (s *sqlStore) FindFood(ctx context.Context, id int) (*foodmodel.Food, error) {
	db := s.db
	var food foodmodel.Food

	if err := db.Where("id = ?", id).
		First(&food).Error; err != nil {
		return nil, err
	}

	return &food, nil
}
