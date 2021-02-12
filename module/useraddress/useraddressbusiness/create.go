package useraddressbusiness

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/useraddress/useraddressmodel"
)

// CreateUserAddressStorage create
type CreateUserAddressStorage interface {
	FindUserAddress(
		ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string) (*useraddressmodel.UserAddress, error)
	CreateUserAddress(ctx context.Context, data *useraddressmodel.UserAddressCreate) error
}

type createUserAddress struct {
	store CreateUserAddressStorage
}

// NewCreateUserAddressBiz create
func NewCreateUserAddressBiz(store CreateUserAddressStorage) *createUserAddress {
	return &createUserAddress{store: store}
}

func (biz *createUserAddress) CreateNewUserAddress(ctx context.Context, data *useraddressmodel.UserAddressCreate) error {
	userAddress, err := biz.store.FindUserAddress(ctx, map[string]interface{}{"title": data.Title})

	if userAddress != nil {
		return common.ErrEntityExisted(useraddressmodel.EntityName, err)
	}

	if err := biz.store.CreateUserAddress(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(useraddressmodel.EntityName, err)
	}

	return nil
}
