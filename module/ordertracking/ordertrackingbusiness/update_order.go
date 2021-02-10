package ordertrackingbusiness

import (
	"context"
	"errors"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/ordertracking/ordertrackingmodel"
)

// UpdateOrderTrackingStorage update
type UpdateOrderTrackingStorage interface {
	FindOrderTracking(
		ctx context.Context,
		id int,
		moreInfo ...string) (*ordertrackingmodel.OrderTracking, error)
	UpdateOrderTracking(ctx context.Context, id int, data *ordertrackingmodel.OrderTrackingUpdate) error
}

type updateOrderTracking struct {
	store UpdateOrderTrackingStorage
}

// NewUpdateOrderTrackingBiz update
func NewUpdateOrderTrackingBiz(store UpdateOrderTrackingStorage) *updateOrderTracking {
	return &updateOrderTracking{store: store}
}

func (biz *updateOrderTracking) UpdateOrderTracking(ctx context.Context, id int, data *ordertrackingmodel.OrderTrackingUpdate) error {
	ordertracking, err := biz.store.FindOrderTracking(ctx, id)
	if err != nil {
		return common.ErrCannotGetEntity(ordertrackingmodel.EntityName, err)
	}

	if ordertracking.Status == 0 {
		return common.ErrCannotGetEntity(ordertrackingmodel.EntityName, errors.New("ordertracking not found"))
	}

	if err := biz.store.UpdateOrderTracking(ctx, ordertracking.ID, data); err != nil {
		return common.ErrCannotUpdateEntity(ordertrackingmodel.EntityName, err)
	}

	return nil
}
