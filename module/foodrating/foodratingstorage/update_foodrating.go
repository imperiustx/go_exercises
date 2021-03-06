package foodratingstorage

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/foodrating/foodratingmodel"
)

func (s *sqlStore) UpdateFoodRating(
	ctx context.Context,
	conditions map[string]interface{},
	data *foodratingmodel.FoodRatingUpdate) error {

	db := s.db.Table(data.TableName())

	if err := db.Where(conditions).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
