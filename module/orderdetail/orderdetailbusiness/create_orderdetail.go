package orderdetailbusiness

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/orderdetail/orderdetailmodel"
)

// CreateOrderDetailStorage create
type CreateOrderDetailStorage interface {
	CreateOrderDetail(ctx context.Context, data *orderdetailmodel.OrderDetailCreate) error
}

type createOrderDetail struct {
	store CreateOrderDetailStorage
}

// NewCreateOrderDetailBiz create
func NewCreateOrderDetailBiz(store CreateOrderDetailStorage) *createOrderDetail {
	return &createOrderDetail{store: store}
}

func (biz *createOrderDetail) CreateNewOrderDetail(ctx context.Context, data *orderdetailmodel.OrderDetailCreate) error {
	if err := biz.store.CreateOrderDetail(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(orderdetailmodel.EntityName, err)
	}

	return nil
}
