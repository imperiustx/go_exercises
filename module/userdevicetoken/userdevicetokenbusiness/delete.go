package userdevicetokenbusiness

import (
	"context"
	"errors"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/userdevicetoken/userdevicetokenmodel"
)

// DeleteUserDeviceTokenStorage delete
type DeleteUserDeviceTokenStorage interface {
	FindUserDeviceToken(ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string) (*userdevicetokenmodel.UserDeviceToken, error)
	DeleteUserDeviceToken(
		ctx context.Context,
		conditions map[string]interface{}) error
}

type deleteUserDeviceToken struct {
	store DeleteUserDeviceTokenStorage
}

// NewDeleteUserDeviceTokenBiz delete
func NewDeleteUserDeviceTokenBiz(store DeleteUserDeviceTokenStorage) *deleteUserDeviceToken {
	return &deleteUserDeviceToken{
		store: store,
	}
}

func (biz *deleteUserDeviceToken) DeleteUserDeviceToken(
	ctx context.Context,
	conditions map[string]interface{}) error {

	userdevicetoken, err := biz.store.FindUserDeviceToken(ctx, conditions)
	if err != nil {
		return common.ErrCannotGetEntity(userdevicetokenmodel.EntityName, err)
	}

	if userdevicetoken.Status == 0 {
		return common.ErrCannotGetEntity(
			userdevicetokenmodel.EntityName,
			errors.New("userdevicetoken not found"),
		)
	}

	if err := biz.store.DeleteUserDeviceToken(ctx, conditions); err != nil {
		return common.ErrCannotDeleteEntity(userdevicetokenmodel.EntityName, err)
	}

	return nil
}
