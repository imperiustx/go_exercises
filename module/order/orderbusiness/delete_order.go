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
		conditions map[string]interface{},
		moreInfo ...string) (*ordermodel.Order, error)
	DeleteOrder(ctx context.Context, conditions map[string]interface{}) error
}

type deleteOrder struct {
	store DeleteOrderStorage
}

// NewDeleteOrderBiz delete
func NewDeleteOrderBiz(store DeleteOrderStorage) *deleteOrder {
	return &deleteOrder{store: store}
}

func (biz *deleteOrder) DeleteOrder(ctx context.Context, conditions map[string]interface{}) error {
	order, err := biz.store.FindOrder(ctx, conditions)
	if err != nil {
		return common.ErrCannotGetEntity(ordermodel.EntityName, err)
	}

	if order.Status == 0 {
		return common.ErrCannotGetEntity(
			ordermodel.EntityName,
			errors.New("order not found"),
		)
	}

	if err := biz.store.DeleteOrder(ctx, conditions); err != nil {
		return err
	}

	return nil
}
