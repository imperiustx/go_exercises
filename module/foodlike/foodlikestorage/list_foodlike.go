package foodlikestorage

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/foodlike/foodlikemodel"
)

func (s *sqlStore) ListFoodLike(ctx context.Context, paging *common.Paging) ([]foodlikemodel.FoodLike, error) {
	db := s.db
	var foodlikes []foodlikemodel.FoodLike

	db = db.Table(foodlikemodel.FoodLike{}.TableName())

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	db = db.Limit(paging.Limit)

	if paging.Cursor > 0 {
		db = db.Where("id < ?", paging.Cursor)
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	if err := db.Find(&foodlikes).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return foodlikes, nil
}

// TODO: list all users liked this food
