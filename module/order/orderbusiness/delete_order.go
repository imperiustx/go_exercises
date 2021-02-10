package orderbusiness

import (
	"context"
	"errors"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/order/ordermodel"
)

// DeleteOrderStorage delete
type DeleteOrderStorage interface {
	FindOrder(
		ctx context.Context,
		id int,
		moreInfo ...string) (*ordermodel.Order, error)
	DeleteOrder(id int) error
}

type deleteOrder struct {
	store DeleteOrderStorage
}

// NewDeleteOrderBiz delete
func NewDeleteOrderBiz(store DeleteOrderStorage) *deleteOrder {
	return &deleteOrder{store: store}
}

func (biz *deleteOrder) DeleteOrder(ctx context.Context, id int) error {
	order, err := biz.store.FindOrder(ctx, id)
	if err != nil {
		return common.ErrCannotGetEntity(ordermodel.EntityName, err)
	}

	if order.Status == 0 {
		return common.ErrCannotGetEntity(
			ordermodel.EntityName,
			errors.New("order not found"),
		)
	}

	if err := biz.store.DeleteOrder(id); err != nil {
		return err
	}

	return nil
}
