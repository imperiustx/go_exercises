package foodlikestorage

import (
	"context"

	"github.com/imperiustx/go_excercises/module/foodlike/foodlikemodel"
)

func (s *sqlStore) FindFoodLike(ctx context.Context, uid, fid int) (*foodlikemodel.FoodLike, error) {
	db := s.db
	var foodlike foodlikemodel.FoodLike

	if err := db.Where(&foodlikemodel.FoodLike{
		UserID: uid,
		FoodID: fid,
	}).First(&foodlike).Error; err != nil {
		return nil, err
	}

	return &foodlike, nil
}
