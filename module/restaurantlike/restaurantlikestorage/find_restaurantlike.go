package restaurantlikestorage

import (
	"context"

	"github.com/imperiustx/go_excercises/module/restaurantlike/restaurantlikemodel"
)

func (s *sqlStore) FindRestaurantLike(
	ctx context.Context,
	uid, rid int,
	moreInfo ...string) (*restaurantlikemodel.RestaurantLike, error) {

	db := s.db.Table(restaurantlikemodel.RestaurantLike{}.TableName())

	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}

	var restaurantlike restaurantlikemodel.RestaurantLike

	if err := db.Where(&restaurantlikemodel.RestaurantLike{
		UserID:       uid,
		RestaurantID: rid,
	}).First(&restaurantlike).Error; err != nil {
		return nil, err
	}

	return &restaurantlike, nil
}
