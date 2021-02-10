package cartstorage

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/cart/cartmodel"
)

func (s *sqlStore) UpdateCart(ctx context.Context, uid, fid int, data *cartmodel.CartUpdate) error {
	db := s.db.Table(cartmodel.Cart{}.TableName())

	if err := db.Where(&cartmodel.Cart{
		UserID: uid,
		FoodID: fid,
	}).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
