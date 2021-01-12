package addressbusiness

import "github.com/imperiustx/go_excercises/module/address/addressmodel"

// ListAddressStorage list
type ListAddressStorage interface {
	ListAddress() ([]addressmodel.Address, error)
}

type listAddress struct {
	store ListAddressStorage
}

// NewListAddressBiz list
func NewListAddressBiz(store ListAddressStorage) *listAddress {
	return &listAddress{store: store}
}

func (biz *listAddress) ListAllAddress() ([]addressmodel.Address, error) {
	return biz.store.ListAddress()
}
