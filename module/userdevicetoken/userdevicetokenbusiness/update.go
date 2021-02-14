package userdevicetokenbusiness

import (
	"context"
	"errors"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/userdevicetoken/userdevicetokenmodel"
)

// UpdateUserDeviceTokenStorage update
type UpdateUserDeviceTokenStorage interface {
	FindUserDeviceToken(
		ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string) (*userdevicetokenmodel.UserDeviceToken, error)
	UpdateUserDeviceToken(
		ctx context.Context,
		conditions map[string]interface{},
		data *userdevicetokenmodel.UserDeviceTokenUpdate) error
}

type updateUserDeviceToken struct {
	store     UpdateUserDeviceTokenStorage
}

// NewUpdateUserDeviceTokenBiz update
func NewUpdateUserDeviceTokenBiz(store UpdateUserDeviceTokenStorage) *updateUserDeviceToken {
	return &updateUserDeviceToken{
		store:     store,
	}
}

func (biz *updateUserDeviceToken) UpdateUserDeviceToken(
	ctx context.Context,
	conditions map[string]interface{},
	data *userdevicetokenmodel.UserDeviceTokenUpdate) error {

	userdevicetoken, err := biz.store.FindUserDeviceToken(ctx, conditions)
	if err != nil {
		return common.ErrCannotGetEntity(userdevicetokenmodel.EntityName, err)
	}

	if userdevicetoken.Status == 0 {
		return common.ErrCannotGetEntity(
			userdevicetokenmodel.EntityName, 
			errors.New("userdevicetoken not found"))
	}

	if err := biz.store.UpdateUserDeviceToken(ctx, conditions, data); err != nil {
		return common.ErrCannotUpdateEntity(userdevicetokenmodel.EntityName, err)
	}

	return nil
}
