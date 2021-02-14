package foodstorage

import (
	"context"

	"github.com/imperiustx/go_excercises/module/food/foodmodel"
)

func (s *sqlStore) FindFood(
	ctx context.Context, 
	conditions map[string]interface{}, 
	moreInfo ...string) (*foodmodel.Food, error) {

	db := s.db.Table(foodmodel.Food{}.TableName())

	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}

	var food foodmodel.Food

	if err := db.Where(conditions).First(&food).Error; err != nil {
		return nil, err
	}

	return &food, nil
}
