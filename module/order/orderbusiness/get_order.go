package orderbusiness

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/order/ordermodel"
)

// GetOrderStorage get
type GetOrderStorage interface {
	FindOrder(
		ctx context.Context, 
		id int,
		moreInfo ...string) (*ordermodel.Order, error)
}

type getOrderBiz struct {
	store GetOrderStorage
}

// NewGetOrderBiz get
func NewGetOrderBiz(store GetOrderStorage) *getOrderBiz {
	return &getOrderBiz{store: store}
}

func (biz *getOrderBiz) GetOrder(
	ctx context.Context, 
	id int,
	moreInfo ...string) (*ordermodel.Order, error) {
	order, err := biz.store.FindOrder(ctx, id)
	if err != nil {
		return nil, common.ErrCannotGetEntity(ordermodel.EntityName, err)
	}

	return order, nil
}
