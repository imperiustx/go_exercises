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
		conditions map[string]interface{},
		moreInfo ...string) (*ordertrackingmodel.OrderTracking, error)
	DeleteOrderTracking(
		ctx context.Context,
		conditions map[string]interface{}) error
}

type deleteOrderTracking struct {
	store DeleteOrderTrackingStorage
}

// NewDeleteOrderTrackingBiz delete
func NewDeleteOrderTrackingBiz(store DeleteOrderTrackingStorage) *deleteOrderTracking {
	return &deleteOrderTracking{store: store}
}

func (biz *deleteOrderTracking) DeleteOrderTracking(ctx context.Context, conditions map[string]interface{}) error {
	ordertracking, err := biz.store.FindOrderTracking(ctx, conditions)
	if err != nil {
		return common.ErrCannotGetEntity(ordertrackingmodel.EntityName, err)
	}

	if ordertracking.Status == 0 {
		return common.ErrCannotGetEntity(
			ordertrackingmodel.EntityName,
			errors.New("ordertracking not found"),
		)
	}

	if err := biz.store.DeleteOrderTracking(ctx, conditions); err != nil {
		return err
	}

	return nil
}
