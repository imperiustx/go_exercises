package useraddressbusiness

import "github.com/imperiustx/go_excercises/module/useraddress/useraddressmodel"

// ListUserAddressStorage list
type ListUserAddressStorage interface {
	ListUserAddress() ([]useraddressmodel.UserAddress, error)
}

type listUserAddress struct {
	store ListUserAddressStorage
}

// NewListUserAddressBiz list
func NewListUserAddressBiz(store ListUserAddressStorage) *listUserAddress {
	return &listUserAddress{store: store}
}

func (biz *listUserAddress) ListAllUserAddress() ([]useraddressmodel.UserAddress, error) {
	return biz.store.ListUserAddress()
}
