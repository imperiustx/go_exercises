package foodstorage

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/food/foodmodel"
)

func (s *sqlStore) UpdateFood(
	ctx context.Context,
	conditions map[string]interface{},
	data *foodmodel.FoodUpdate) error {

	db := s.db.Table(data.TableName())

	if err := db.Where(conditions).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
