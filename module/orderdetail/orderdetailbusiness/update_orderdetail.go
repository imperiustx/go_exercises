package orderdetailbusiness

import (
	"context"
	"errors"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/orderdetail/orderdetailmodel"
)

// UpdateOrderDetailStorage update
type UpdateOrderDetailStorage interface {
	FindOrderDetail(ctx context.Context, id int) (*orderdetailmodel.OrderDetail, error)
	UpdateOrderDetail(ctx context.Context, id int, data *orderdetailmodel.OrderDetailUpdate) error
}

type updateOrderDetail struct {
	store UpdateOrderDetailStorage
}

// NewUpdateOrderDetailBiz update
func NewUpdateOrderDetailBiz(store UpdateOrderDetailStorage) *updateOrderDetail {
	return &updateOrderDetail{store: store}
}

func (biz *updateOrderDetail) UpdateOrderDetail(ctx context.Context, id int, data *orderdetailmodel.OrderDetailUpdate) error {
	orderdetail, err := biz.store.FindOrderDetail(ctx, id)
	if err != nil {
		return common.ErrCannotGetEntity(orderdetailmodel.EntityName, err)
	}

	if orderdetail.Status == 0 {
		return common.ErrCannotGetEntity(orderdetailmodel.EntityName, errors.New("orderdetail not found"))
	}

	if err := biz.store.UpdateOrderDetail(ctx, orderdetail.ID, data); err != nil {
		return common.ErrCannotUpdateEntity(orderdetailmodel.EntityName, err)
	}

	return nil
}
