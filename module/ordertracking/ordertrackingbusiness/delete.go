package ordertrackingbusiness

import (
	"context"
	"errors"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/ordertracking/ordertrackingmodel"
)

// DeleteOrderTrackingStorage delete
type DeleteOrderTrackingStorage interface {
	FindOrderTracking(
		ctx context.Context,
		id int,
		moreInfo ...string) (*ordertrackingmodel.OrderTracking, error)
	DeleteOrderTracking(id int) error
}

type deleteOrderTracking struct {
	store DeleteOrderTrackingStorage
}

// NewDeleteOrderTrackingBiz delete
func NewDeleteOrderTrackingBiz(store DeleteOrderTrackingStorage) *deleteOrderTracking {
	return &deleteOrderTracking{store: store}
}

func (biz *deleteOrderTracking) DeleteOrderTracking(ctx context.Context, id int) error {
	ordertracking, err := biz.store.FindOrderTracking(ctx, id)
	if err != nil {
		return common.ErrCannotGetEntity(ordertrackingmodel.EntityName, err)
	}

	if ordertracking.Status == 0 {
		return common.ErrCannotGetEntity(
			ordertrackingmodel.EntityName,
			errors.New("ordertracking not found"),
		)
	}

	if err := biz.store.DeleteOrderTracking(id); err != nil {
		return err
	}

	return nil
}
