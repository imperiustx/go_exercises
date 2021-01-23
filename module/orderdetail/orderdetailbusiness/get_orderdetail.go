package orderdetailbusiness

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/orderdetail/orderdetailmodel"
)

// GetOrderDetailStorage get
type GetOrderDetailStorage interface {
	FindOrderDetail(ctx context.Context, id int) (*orderdetailmodel.OrderDetail, error)
}

type getOrderDetailBiz struct {
	store GetOrderDetailStorage
}

// NewGetOrderDetailBiz get
func NewGetOrderDetailBiz(store GetOrderDetailStorage) *getOrderDetailBiz {
	return &getOrderDetailBiz{store: store}
}

func (biz *getOrderDetailBiz) GetOrderDetail(ctx context.Context, id int) (*orderdetailmodel.OrderDetail, error) {
	orderdetail, err := biz.store.FindOrderDetail(ctx, id)
	if err != nil {
		return nil, common.ErrCannotGetEntity(orderdetailmodel.EntityName, err)
	}

	return orderdetail, nil
}
