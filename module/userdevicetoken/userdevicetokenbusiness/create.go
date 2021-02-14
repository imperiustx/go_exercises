package userdevicetokenbusiness

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/userdevicetoken/userdevicetokenmodel"
)

type CreateUserDeviceTokenStorage interface {
	FindUserDeviceTokenID(
		ctx context.Context,
		uid int,
		os string,
		moreInfo ...string) (*int, error)
	CreateUserDeviceToken(
		ctx context.Context, 
		data *userdevicetokenmodel.UserDeviceTokenCreate) error
}

type createUserDeviceTokenBusiness struct {
	createUserDeviceTokenStorage CreateUserDeviceTokenStorage
}

func NewCreateUserDeviceTokenBusiness(createUserDeviceTokenStorage CreateUserDeviceTokenStorage) *createUserDeviceTokenBusiness {
	return &createUserDeviceTokenBusiness{
		createUserDeviceTokenStorage: createUserDeviceTokenStorage,
	}
}

func (business *createUserDeviceTokenBusiness) CreateUserDeviceToken(
	ctx context.Context,
	data *userdevicetokenmodel.UserDeviceTokenCreate) error {

	userdevicetoken, err := business.createUserDeviceTokenStorage.FindUserDeviceTokenID(
		ctx,
		data.UserID,
		data.OS,
	)

	if userdevicetoken != nil {
		return common.ErrEntityExisted(userdevicetokenmodel.EntityName, err)
	}

	if err := business.createUserDeviceTokenStorage.CreateUserDeviceToken(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(userdevicetokenmodel.EntityName, err)
	}

	return nil
}
