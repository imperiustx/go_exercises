package userdevicetokenbusiness

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/userdevicetoken/userdevicetokenmodel"
)

// ListUserDeviceTokenStorage list
type ListUserDeviceTokenStorage interface {
	ListUserDeviceToken(
		ctx context.Context,
		filter *userdevicetokenmodel.Filter,
		paging *common.Paging,
		order *common.OrderSort,
		moreKeys ...string) ([]userdevicetokenmodel.UserDeviceToken, error)
}

type listUserDeviceToken struct {
	store     ListUserDeviceTokenStorage
}

// NewListUserDeviceTokenBiz list
func NewListUserDeviceTokenBiz(store ListUserDeviceTokenStorage) *listUserDeviceToken {
	return &listUserDeviceToken{
		store:     store,
	}
}

func (biz *listUserDeviceToken) ListAllUserDeviceToken(
	ctx context.Context,
	filter *userdevicetokenmodel.Filter,
	paging *common.Paging,
	order *common.OrderSort) ([]userdevicetokenmodel.UserDeviceToken, error) {

	data, err := biz.store.ListUserDeviceToken(ctx, filter, paging, order)
	if err != nil {
		return nil, common.ErrCannotListEntity(userdevicetokenmodel.EntityName, err)
	}

	return data, nil
}
