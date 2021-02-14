package foodratingstorage

import (
	"context"

	"github.com/imperiustx/go_excercises/module/foodrating/foodratingmodel"
)

func (s *sqlStore) FindFoodRatingID(
	ctx context.Context,
	uid, fid int,
	moreInfo ...string) (*int, error) {

	db := s.db.Table(foodratingmodel.FoodRating{}.TableName())
	var foodrating foodratingmodel.FoodRating

	if err := db.Where(&foodratingmodel.FoodRating{
		UserID: uid,
		FoodID: fid,
	}).First(&foodrating).Error; err != nil {
		return nil, err
	}

	return &foodrating.ID, nil
}
