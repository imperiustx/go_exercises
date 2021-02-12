package useraddressbusiness

import (
	"context"
	"errors"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/useraddress/useraddressmodel"
)

// UpdateUserAddressStorage update
type UpdateUserAddressStorage interface {
	FindUserAddress(
		ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string) (*useraddressmodel.UserAddress, error)
	UpdateUserAddress(
		ctx context.Context, 
		conditions map[string]interface{},
		data *useraddressmodel.UserAddressUpdate) error
}

type updateUserAddress struct {
	store UpdateUserAddressStorage
}

// NewUpdateUserAddressBiz update
func NewUpdateUserAddressBiz(store UpdateUserAddressStorage) *updateUserAddress {
	return &updateUserAddress{store: store}
}

func (biz *updateUserAddress) UpdateUserAddress(
	ctx context.Context, 
	conditions map[string]interface{}, 
	data *useraddressmodel.UserAddressUpdate) error {

	useraddress, err := biz.store.FindUserAddress(ctx, conditions)
	if err != nil {
		return common.ErrCannotGetEntity(useraddressmodel.EntityName, err)
	}

	if useraddress.Status == 0 {
		return common.ErrCannotGetEntity(useraddressmodel.EntityName, errors.New("useraddress not found"))
	}

	if err := biz.store.UpdateUserAddress(ctx, conditions, data); err != nil {
		return common.ErrCannotUpdateEntity(useraddressmodel.EntityName, err)
	}

	return nil
}

// TODO: add requester