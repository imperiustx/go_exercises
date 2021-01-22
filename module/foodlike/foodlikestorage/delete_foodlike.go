package foodlikestorage

import "github.com/imperiustx/go_excercises/module/foodlike/foodlikemodel"

func (s *sqlStore) DeleteFoodLike(uid, fid int) error {
	db := s.db

	// TODO: recheck this one
	if err := db.Table("food_likes").
		Delete(&foodlikemodel.FoodLike{
			UserID: uid,
			FoodID: fid,
		}).Error; err != nil {
		return err
	}

	return nil
}
