package cartbusiness

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/cart/cartmodel"
)

type GetCartStorage interface {
	FindCart(
		ctx context.Context,
		uid, fid int,
		moreInfo ...string) (*cartmodel.Cart, error)
}

type getCartBiz struct {
	store GetCartStorage
}

func NewGetCartBiz(store GetCartStorage) *getCartBiz {
	return &getCartBiz{store: store}
}

func (biz *getCartBiz) GetCart(ctx context.Context, uid, fid int, moreInfo ...string) (*cartmodel.Cart, error) {

	cart, err := biz.store.FindCart(ctx, uid, fid)
	if err != nil {
		return nil, common.ErrCannotGetEntity(cartmodel.EntityName, err)
	}

	return cart, nil
}

//TODO: get cart by user_id and get cart by cart_id