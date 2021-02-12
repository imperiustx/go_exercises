package useraddressbusiness

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/useraddress/useraddressmodel"
)

// GetUserAddressStorage get
type GetUserAddressStorage interface {
	FindUserAddress(
		ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string) (*useraddressmodel.UserAddress, error)
}

type getUserAddressBiz struct {
	store GetUserAddressStorage
}

// NewGetUserAddressBiz get
func NewGetUserAddressBiz(store GetUserAddressStorage) *getUserAddressBiz {
	return &getUserAddressBiz{store: store}
}

func (biz *getUserAddressBiz) GetUserAddress(
	ctx context.Context,
	conditions map[string]interface{},
	moreInfo ...string) (*useraddressmodel.UserAddress, error) {

	useraddress, err := biz.store.FindUserAddress(ctx, conditions)
	if err != nil {
		return nil, common.ErrCannotGetEntity(useraddressmodel.EntityName, err)
	}

	return useraddress, nil
}
