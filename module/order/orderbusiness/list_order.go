package orderbusiness

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/order/ordermodel"
)

// ListOrderStorage list
type ListOrderStorage interface {
	ListOrder(
		ctx context.Context,
		paging *common.Paging,
		order *common.OrderSort,
		moreKeys ...string) ([]ordermodel.Order, error)
}

type listOrder struct {
	store     ListOrderStorage
	// requester common.Requester
}

// NewListOrderBiz list
func NewListOrderBiz(store ListOrderStorage) *listOrder {
	return &listOrder{
		store:     store,
		// requester: requester,
	}
}

func (biz *listOrder) ListAllOrder(
	ctx context.Context,
	paging *common.Paging,
	order *common.OrderSort) ([]ordermodel.Order, error) {

	// if biz.requester.GetRole() != "admin" {
	// 	return []ordermodel.Order{}, common.ErrNoPermission(nil)
	// }

	data, err := biz.store.ListOrder(ctx, paging, order)
	if err != nil {
		return nil, common.ErrCannotListEntity(ordermodel.EntityName, err)
	}

	return data, nil
}
