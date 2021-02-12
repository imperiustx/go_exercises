package useraddressbusiness

import (
	"context"
	"errors"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/useraddress/useraddressmodel"
)

// DeleteUserAddressStorage delete
type DeleteUserAddressStorage interface {
	FindUserAddress(
		ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string) (*useraddressmodel.UserAddress, error)
	DeleteUserAddress(conditions map[string]interface{}) error
}

type deleteUserAddress struct {
	store DeleteUserAddressStorage
}

// NewDeleteUserAddressBiz delete
func NewDeleteUserAddressBiz(store DeleteUserAddressStorage) *deleteUserAddress {
	return &deleteUserAddress{store: store}
}

func (biz *deleteUserAddress) DeleteUserAddress(ctx context.Context, conditions map[string]interface{}) error {
	useraddress, err := biz.store.FindUserAddress(ctx, conditions)
	if err != nil {
		return common.ErrCannotGetEntity(useraddressmodel.EntityName, err)
	}

	if useraddress.Status == 0 {
		return common.ErrCannotGetEntity(useraddressmodel.EntityName, errors.New("useraddress not found"))
	}

	if err := biz.store.DeleteUserAddress(conditions); err != nil {
		return err
	}

	return nil
}

// TODO: add requester