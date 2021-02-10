package cartbusiness

import (
	"context"
	"errors"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/cart/cartmodel"
)

// UpdateCartStorage update
type UpdateCartStorage interface {
	FindCart(
		ctx context.Context,
		uid, fid int,
		moreInfo ...string) (*cartmodel.Cart, error)
	UpdateCart(ctx context.Context, uid, fid int, data *cartmodel.CartUpdate) error
}

type updateCart struct {
	store     UpdateCartStorage
	// requester common.Requester
}

// NewUpdateCartBiz update
func NewUpdateCartBiz(store UpdateCartStorage) *updateCart {
	return &updateCart{
		store:     store,
		// requester: requester,
	}
}

func (biz *updateCart) UpdateCart(ctx context.Context, uid, fid int, data *cartmodel.CartUpdate) error {
	cart, err := biz.store.FindCart(ctx, uid, fid)
	if err != nil {
		return common.ErrCannotGetEntity(cartmodel.EntityName, err)
	}

	if cart.Status == 0 {
		return common.ErrCannotGetEntity(cartmodel.EntityName, errors.New("cart not found"))
	}

	// if biz.requester.GetCartId() != cart.ID {
	// 	return common.ErrNoPermission(nil)
	// }

	if err := biz.store.UpdateCart(ctx, uid, fid, data); err != nil {
		return common.ErrCannotUpdateEntity(cartmodel.EntityName, err)
	}

	return nil
}
