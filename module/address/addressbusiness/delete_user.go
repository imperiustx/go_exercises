package addressbusiness

// DeleteAddressStorage delete
type DeleteAddressStorage interface {
	DeleteAddress(id string) error
}

type deleteAddress struct {
	store DeleteAddressStorage
}

// NewDeleteAddressBiz delete
func NewDeleteAddressBiz(store DeleteAddressStorage) *deleteAddress {
	return &deleteAddress{store: store}
}

func (biz *deleteAddress) DeleteAddress(id string) error {
	return biz.store.DeleteAddress(id)
}
