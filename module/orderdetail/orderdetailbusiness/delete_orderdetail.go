package orderdetailbusiness

import (
	"context"
	"errors"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/orderdetail/orderdetailmodel"
)

// DeleteOrderDetailStorage delete
type DeleteOrderDetailStorage interface {
	FindOrderDetail(ctx context.Context, id int) (*orderdetailmodel.OrderDetail, error)
	DeleteOrderDetail(id int) error
}

type deleteOrderDetail struct {
	store DeleteOrderDetailStorage
}

// NewDeleteOrderDetailBiz delete
func NewDeleteOrderDetailBiz(store DeleteOrderDetailStorage) *deleteOrderDetail {
	return &deleteOrderDetail{store: store}
}

func (biz *deleteOrderDetail) DeleteOrderDetail(ctx context.Context, id int) error {
	orderdetail, err := biz.store.FindOrderDetail(ctx, id)
	if err != nil {
		return common.ErrCannotGetEntity(orderdetailmodel.EntityName, err)
	}

	if orderdetail.Status == 0 {
		return common.ErrCannotGetEntity(orderdetailmodel.EntityName, errors.New("orderdetail not found"))
	}

	if err := biz.store.DeleteOrderDetail(id); err != nil {
		return err
	}

	return nil
}
