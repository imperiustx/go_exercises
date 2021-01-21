package categoryrestaurantstorage

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/categoryrestaurant/categoryrestaurantmodel"
)

func (s *sqlStore) CreateCategoryRestaurant(ctx context.Context, data *categoryrestaurantmodel.CategoryRestaurantCreate) error {
	db := s.db

	if err := db.Table(data.TableName()).
		Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
