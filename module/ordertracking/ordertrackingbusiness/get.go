package ordertrackingbusiness

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/ordertracking/ordertrackingmodel"
)

// GetOrderTrackingStorage get
type GetOrderTrackingStorage interface {
	FindOrderTracking(
		ctx context.Context,
		id int,
		moreInfo ...string) (*ordertrackingmodel.OrderTracking, error)
}

type getOrderTrackingBiz struct {
	store GetOrderTrackingStorage
}

// NewGetOrderTrackingBiz get
func NewGetOrderTrackingBiz(store GetOrderTrackingStorage) *getOrderTrackingBiz {
	return &getOrderTrackingBiz{store: store}
}

func (biz *getOrderTrackingBiz) GetOrderTracking(
	ctx context.Context,
	id int,
	moreInfo ...string) (*ordertrackingmodel.OrderTracking, error) {
	ordertracking, err := biz.store.FindOrderTracking(ctx, id)
	if err != nil {
		return nil, common.ErrCannotGetEntity(ordertrackingmodel.EntityName, err)
	}

	return ordertracking, nil
}
