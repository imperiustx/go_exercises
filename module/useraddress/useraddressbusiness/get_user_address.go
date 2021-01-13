package useraddressbusiness

import "github.com/imperiustx/go_excercises/module/useraddress/useraddressmodel"

// GetUserAddressStorage get
type GetUserAddressStorage interface {
	GetUserAddress(id string) (useraddressmodel.UserAddress, error)
}

type getUserAddress struct {
	store GetUserAddressStorage
}

// NewGetUserAddressBiz get
func NewGetUserAddressBiz(store GetUserAddressStorage) *getUserAddress {
	return &getUserAddress{store: store}
}

func (biz *getUserAddress) GetUserAddress(id string) (useraddressmodel.UserAddress, error) {
	return biz.store.GetUserAddress(id)
}
