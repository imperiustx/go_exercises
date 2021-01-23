package orderdetailbusiness

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/orderdetail/orderdetailmodel"
)

// ListOrderDetailStorage list
type ListOrderDetailStorage interface {
	ListOrderDetail(ctx context.Context, paging *common.Paging) ([]orderdetailmodel.OrderDetail, error)
}

type listOrderDetail struct {
	store ListOrderDetailStorage
}

// NewListOrderDetailBiz list
func NewListOrderDetailBiz(store ListOrderDetailStorage) *listOrderDetail {
	return &listOrderDetail{store: store}
}

func (biz *listOrderDetail) ListAllOrderDetail(ctx context.Context, paging *common.Paging) ([]orderdetailmodel.OrderDetail, error) {
	data, err := biz.store.ListOrderDetail(ctx, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(orderdetailmodel.EntityName, err)
	}

	return data, nil
}
