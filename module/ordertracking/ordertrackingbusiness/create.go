package ordertrackingbusiness

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/ordertracking/ordertrackingmodel"
)

// CreateOrderTrackingStorage create
type CreateOrderTrackingStorage interface {
	CreateOrderTracking(ctx context.Context, data *ordertrackingmodel.OrderTrackingCreate) error
}

type createOrderTracking struct {
	store CreateOrderTrackingStorage
}

// NewCreateOrderTrackingBiz create
func NewCreateOrderTrackingBiz(store CreateOrderTrackingStorage) *createOrderTracking {
	return &createOrderTracking{store: store}
}

func (biz *createOrderTracking) CreateNewOrderTracking(ctx context.Context, data *ordertrackingmodel.OrderTrackingCreate) error {
	if err := biz.store.CreateOrderTracking(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(ordertrackingmodel.EntityName, err)
	}

	return nil
}
