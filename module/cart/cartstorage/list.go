package cartstorage

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/cart/cartmodel"
)

func (s *sqlStore) ListCart(
	ctx context.Context,
	paging *common.Paging,
	moreKeys ...string) ([]cartmodel.Cart, error) {
	db := s.db
	var carts []cartmodel.Cart

	db = db.Table(cartmodel.Cart{}.TableName()).Where("status not in (0)")

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	db = db.Limit(paging.Limit)

	for _, k := range moreKeys {
		db = db.Preload(k)
	}

	if paging.Cursor > 0 {
		db = db.Where("id < ?", paging.Cursor) //TODO: consider this
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	// if o := order; o != nil {
	// 	if o.Order == "asc" {
	// 		db = db.Order("id asc") //FIXME: fix this
	// 	}
	// 	if o.Order == "desc" {
	// 		db = db.Order("id desc") //FIXME: fix this
	// 	}
	// }

	if err := db.Find(&carts).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return carts, nil
}
