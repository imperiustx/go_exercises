package addressbusiness

import "github.com/imperiustx/go_excercises/module/address/addressmodel"

// GetAddressStorage get
type GetAddressStorage interface {
	GetAddress(id string) (addressmodel.Address, error)
}

type getAddress struct {
	store GetAddressStorage
}

// NewGetAddressBiz get
func NewGetAddressBiz(store GetAddressStorage) *getAddress {
	return &getAddress{store: store}
}

func (biz *getAddress) GetAddress(id string) (addressmodel.Address, error) {
	return biz.store.GetAddress(id)
}
