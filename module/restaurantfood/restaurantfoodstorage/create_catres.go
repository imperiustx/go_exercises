package restaurantfoodstorage

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/restaurantfood/restaurantfoodmodel"
)

func (s *sqlStore) CreateRestaurantFood(ctx context.Context, data *restaurantfoodmodel.RestaurantFoodCreate) error {
	db := s.db.Table(data.TableName())

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
