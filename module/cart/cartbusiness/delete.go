package cartbusiness

import (
	"context"
	"errors"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/cart/cartmodel"
)

// DeleteCartStorage delete
type DeleteCartStorage interface {
	FindCart(
		ctx context.Context, 
		uid, fid int, 
		moreInfo ...string) (*cartmodel.Cart, error)
	DeleteCart(uid, fid int) error
}

type deleteCart struct {
	store     DeleteCartStorage
	// requester common.Requester
}

// NewDeleteCartBiz delete
func NewDeleteCartBiz(store DeleteCartStorage) *deleteCart {
	return &deleteCart{
		store:     store,
		// requester: requester,
	}
}

func (biz *deleteCart) DeleteCart(ctx context.Context, uid, fid int, moreInfo ...string) error {
	cart, err := biz.store.FindCart(ctx, uid, fid)
	if err != nil {
		return common.ErrCannotGetEntity(cartmodel.EntityName, err)
	}

	if cart.Status == 0 {
		return common.ErrCannotGetEntity(cartmodel.EntityName, errors.New("cart not found"))
	}

	// isAuthor := biz.requester.GetCartId() == cart.ID 
	// TODO: consider this logic
	// isAdmin := biz.requester.GetRole() == "admin"

	// if !isAdmin {
	// 	return common.ErrNoPermission(nil)
	// }

	if err := biz.store.DeleteCart(uid, fid); err != nil {
		return common.ErrCannotDeleteEntity(cartmodel.EntityName, err)
	}

	return nil
}
