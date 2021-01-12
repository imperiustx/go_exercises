package addressbusiness

// UpdateAddressStorage update
type UpdateAddressStorage interface {
	UpdateAddress(id string, v interface{}) error
}

type updateAddress struct {
	store UpdateAddressStorage
}

// NewUpdateAddressBiz update
func NewUpdateAddressBiz(store UpdateAddressStorage) *updateAddress {
	return &updateAddress{store: store}
}

func (biz *updateAddress) UpdateAddress(id string, v interface{}) error {
	return biz.store.UpdateAddress(id, v)
}
