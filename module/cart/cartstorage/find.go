package cartstorage

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/cart/cartmodel"
)

func (s *sqlStore) FindCart(ctx context.Context, uid, fid int, moreInfo ...string) (*cartmodel.Cart, error) {
	db := s.db.Table(cartmodel.Cart{}.TableName())
	var cart cartmodel.Cart

	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}

	if err := db.Where(&cartmodel.Cart{
		UserID: uid,
		FoodID: fid,
	}).First(&cart).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return &cart, nil
}
