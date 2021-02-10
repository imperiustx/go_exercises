package ordertrackingbusiness

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/ordertracking/ordertrackingmodel"
)

// ListOrderTrackingStorage list
type ListOrderTrackingStorage interface {
	ListOrderTracking(
		ctx context.Context,
		paging *common.Paging,
		ordertracking *common.OrderSort,
		moreKeys ...string) ([]ordertrackingmodel.OrderTracking, error)
}

type listOrderTracking struct {
	store ListOrderTrackingStorage
	// requester common.Requester
}

// NewListOrderTrackingBiz list
func NewListOrderTrackingBiz(store ListOrderTrackingStorage) *listOrderTracking {
	return &listOrderTracking{
		store: store,
		// requester: requester,
	}
}

func (biz *listOrderTracking) ListAllOrderTracking(
	ctx context.Context,
	paging *common.Paging,
	ordertracking *common.OrderSort) ([]ordertrackingmodel.OrderTracking, error) {

	// if biz.requester.GetRole() != "admin" {
	// 	return []ordertrackingmodel.OrderTracking{}, common.ErrNoPermission(nil)
	// }

	data, err := biz.store.ListOrderTracking(ctx, paging, ordertracking)
	if err != nil {
		return nil, common.ErrCannotListEntity(ordertrackingmodel.EntityName, err)
	}

	return data, nil
}
