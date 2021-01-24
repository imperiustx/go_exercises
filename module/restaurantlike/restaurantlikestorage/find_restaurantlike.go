package restaurantlikestorage

import (
	"context"

	"github.com/imperiustx/go_excercises/module/restaurantlike/restaurantlikemodel"
)

func (s *sqlStore) FindRestaurantLike(ctx context.Context, uid, rid int) (*restaurantlikemodel.RestaurantLike, error) {
	db := s.db
	var restaurantlike restaurantlikemodel.RestaurantLike

	if err := db.Where(&restaurantlikemodel.RestaurantLike{
		UserID:       uid,
		RestaurantID: rid,
	}).First(&restaurantlike).Error; err != nil {
		return nil, err
	}

	return &restaurantlike, nil
}
