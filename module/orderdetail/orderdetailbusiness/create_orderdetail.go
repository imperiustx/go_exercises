package orderdetailbusiness

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/orderdetail/orderdetailmodel"
)

// CreateOrderDetailStorage create
type CreateOrderDetailStorage interface {
	FindOrderDetail(
		ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string) (*orderdetailmodel.OrderDetail, error)
	CreateOrderDetail(
		ctx context.Context,
		data *orderdetailmodel.OrderDetailCreate) error
}

type createOrderDetail struct {
	store CreateOrderDetailStorage
}

// NewCreateOrderDetailBiz create
func NewCreateOrderDetailBiz(store CreateOrderDetailStorage) *createOrderDetail {
	return &createOrderDetail{store: store}
}

func (biz *createOrderDetail) CreateNewOrderDetail(
	ctx context.Context,
	data *orderdetailmodel.OrderDetailCreate) error {

	orderDetail, err := biz.store.FindOrderDetail(
		ctx,
		map[string]interface{}{"id": data.ID}, // TODO: consider this logic
	)

	if orderDetail != nil {
		return common.ErrEntityExisted(orderdetailmodel.EntityName, err)
	}

	if err := biz.store.CreateOrderDetail(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(orderdetailmodel.EntityName, err)
	}

	return nil
}
