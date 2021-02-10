package cartstorage

import "github.com/imperiustx/go_excercises/module/cart/cartmodel"

func (s *sqlStore) DeleteCart(uid, fid int) error {
	db := s.db.Table(cartmodel.Cart{}.TableName())

	if err := db.Where(&cartmodel.Cart{
		UserID: uid,
		FoodID: fid,
	}).Update("status", 0).Error; err != nil {
		return err
	}

	return nil
}
