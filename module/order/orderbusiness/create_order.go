package orderbusiness

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/order/ordermodel"
)

// CreateOrderStorage create
type CreateOrderStorage interface {
	CreateOrder(ctx context.Context, data *ordermodel.OrderCreate) error
}

type createOrder struct {
	store CreateOrderStorage
}

// NewCreateOrderBiz create
func NewCreateOrderBiz(store CreateOrderStorage) *createOrder {
	return &createOrder{store: store}
}

func (biz *createOrder) CreateNewOrder(ctx context.Context, data *ordermodel.OrderCreate) error {
	if err := biz.store.CreateOrder(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(ordermodel.EntityName, err)
	}

	return nil
}
