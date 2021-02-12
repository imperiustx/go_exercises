package useraddressbusiness

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/useraddress/useraddressmodel"
)

// ListUserAddressStorage list
type ListUserAddressStorage interface {
	ListUserAddress(
		ctx context.Context,
		filter *useraddressmodel.Filter,
		paging *common.Paging,
		order *common.OrderSort,
		moreKeys ...string) ([]useraddressmodel.UserAddress, error)
}

type listUserAddress struct {
	store ListUserAddressStorage
}

// NewListUserAddressBiz list
func NewListUserAddressBiz(store ListUserAddressStorage) *listUserAddress {
	return &listUserAddress{store: store}
}

func (biz *listUserAddress) ListAllUserAddress(
	ctx context.Context,
	filter *useraddressmodel.Filter,
	paging *common.Paging,
	order *common.OrderSort,
	moreKeys ...string) ([]useraddressmodel.UserAddress, error) {

	data, err := biz.store.ListUserAddress(ctx, filter, paging, order)
	if err != nil {
		return nil, common.ErrCannotListEntity(useraddressmodel.EntityName, err)
	}

	return data, nil
}

// TODO: add requester