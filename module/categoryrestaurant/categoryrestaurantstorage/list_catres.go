package categoryrestaurantstorage

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/categoryrestaurant/categoryrestaurantmodel"
)

func (s *sqlStore) ListCategoryRestaurant(ctx context.Context, paging *common.Paging) ([]categoryrestaurantmodel.CategoryRestaurant, error) {
	db := s.db
	var categoryrestaurants []categoryrestaurantmodel.CategoryRestaurant

	db = db.Table(categoryrestaurantmodel.CategoryRestaurant{}.TableName()).Where("status not in (0)")

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	db = db.Limit(paging.Limit)
	//TODO: fix id
	if paging.Cursor > 0 {
		db = db.Where("id < ?", paging.Cursor)
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	// id desc
	if err := db.Order("id asc").Find(&categoryrestaurants).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return categoryrestaurants, nil
}
