package orderbusiness

import (
	"context"
	"errors"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/order/ordermodel"
)

// UpdateOrderStorage update
type UpdateOrderStorage interface {
	FindOrder(
		ctx context.Context, 
		id int,
		moreInfo ...string) (*ordermodel.Order, error)
	UpdateOrder(ctx context.Context, id int, data *ordermodel.OrderUpdate) error
}

type updateOrder struct {
	store UpdateOrderStorage
}

// NewUpdateOrderBiz update
func NewUpdateOrderBiz(store UpdateOrderStorage) *updateOrder {
	return &updateOrder{store: store}
}

func (biz *updateOrder) UpdateOrder(ctx context.Context, id int, data *ordermodel.OrderUpdate) error {
	order, err := biz.store.FindOrder(ctx, id)
	if err != nil {
		return common.ErrCannotGetEntity(ordermodel.EntityName, err)
	}

	if order.Status == 0 {
		return common.ErrCannotGetEntity(ordermodel.EntityName, errors.New("order not found"))
	}

	if err := biz.store.UpdateOrder(ctx, order.ID, data); err != nil {
		return common.ErrCannotUpdateEntity(ordermodel.EntityName, err)
	}

	return nil
}
