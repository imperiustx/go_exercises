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
}

// NewListCartBiz list
func NewListCartBiz(store ListCartStorage) *listCart {
	return &listCart{
		store:     store,
	}
}

func (biz *listCart) ListAllCart(
	ctx context.Context,
	paging *common.Paging) ([]cartmodel.Cart, error) {

	data, err := biz.store.ListCart(ctx, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(cartmodel.EntityName, err)
	}

	return data, nil
}
