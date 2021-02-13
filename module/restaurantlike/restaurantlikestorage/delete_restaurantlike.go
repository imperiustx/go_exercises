package restaurantlikestorage

import (
	"context"

	"github.com/imperiustx/go_excercises/module/restaurantlike/restaurantlikemodel"
)

func (s *sqlStore) DeleteRestaurantLike(ctx context.Context, uid, rid int) error {
	db := s.db.Table(restaurantlikemodel.RestaurantLike{}.TableName())

	if err := db.Delete(&restaurantlikemodel.RestaurantLike{
		UserID:       uid,
		RestaurantID: rid,
	}).Error; err != nil {
		return err
	}

	return nil
}
