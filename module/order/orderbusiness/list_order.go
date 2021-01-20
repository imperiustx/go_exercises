package orderbusiness

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/order/ordermodel"
)

// ListOrderStorage list
type ListOrderStorage interface {
	ListOrder(ctx context.Context, paging *common.Paging) ([]ordermodel.Order, error)
}

type listOrder struct {
	store ListOrderStorage
}

// NewListOrderBiz list
func NewListOrderBiz(store ListOrderStorage) *listOrder {
	return &listOrder{store: store}
}

func (biz *listOrder) ListAllOrder(ctx context.Context, paging *common.Paging) ([]ordermodel.Order, error) {
	data, err := biz.store.ListOrder(ctx, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(ordermodel.EntityName, err)
	}

	return data, nil
}
