package userdevicetokenbusiness

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/userdevicetoken/userdevicetokenmodel"
)

type GetUserDeviceTokenStorage interface {
	FindUserDeviceToken(
		ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string) (*userdevicetokenmodel.UserDeviceToken, error)
}

type getUserDeviceTokenBiz struct {
	store GetUserDeviceTokenStorage
}

func NewGetUserDeviceTokenBiz(store GetUserDeviceTokenStorage) *getUserDeviceTokenBiz {
	return &getUserDeviceTokenBiz{store: store}
}

func (biz *getUserDeviceTokenBiz) GetUserDeviceToken(
	ctx context.Context,
	conditions map[string]interface{},
	moreInfo ...string) (*userdevicetokenmodel.UserDeviceToken, error) {

	userdevicetoken, err := biz.store.FindUserDeviceToken(ctx, conditions)
	if err != nil {
		return nil, common.ErrCannotGetEntity(userdevicetokenmodel.EntityName, err)
	}

	return userdevicetoken, nil
}
