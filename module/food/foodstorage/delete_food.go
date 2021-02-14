package foodstorage

import (
	"context"

	"github.com/imperiustx/go_excercises/module/food/foodmodel"
)

func (s *sqlStore) DeleteFood(
	ctx context.Context,
	conditions map[string]interface{}) error {

	db := s.db.Table(foodmodel.Food{}.TableName())

	if err := db.Where(conditions).Update("status", 0).Error; err != nil {
		return err
	}

	return nil
}
