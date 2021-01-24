package restaurantlikestorage

import "github.com/imperiustx/go_excercises/module/restaurantlike/restaurantlikemodel"

func (s *sqlStore) DeleteRestaurantLike(uid, rid int) error {
	db := s.db

	// TODO: recheck this one
	if err := db.Table("restaurant_likes").
		Delete(&restaurantlikemodel.RestaurantLike{
			UserID:       uid,
			RestaurantID: rid,
		}).Error; err != nil {
		return err
	}

	return nil
}
