package cartbusiness

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/cart/cartmodel"
)

// ListCartStorage list
type ListCartStorage interface {
	ListCart(
		ctx context.Context,
		paging *common.Paging,
		moreKeys ...string) ([]cartmodel.Cart, error)
}

type listCart struct {
	store     ListCartStorage
	// requester common.Requester
}

// NewListCartBiz list
func NewListCartBiz(store ListCartStorage) *listCart {
	return &listCart{
		store:     store,
		// requester: requester,
	}
}

func (biz *listCart) ListAllCart(
	ctx context.Context,
	paging *common.Paging) ([]cartmodel.Cart, error) {

	// if biz.requester.GetRole() != "admin" {
	// 	return []cartmodel.Cart{}, common.ErrNoPermission(nil)
	// }

	data, err := biz.store.ListCart(ctx, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(cartmodel.EntityName, err)
	}

	return data, nil
}
