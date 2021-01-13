package useraddressbusiness

// UpdateUserAddressStorage update
type UpdateUserAddressStorage interface {
	UpdateUserAddress(id string, v interface{}) error
}

type updateUserAddress struct {
	store UpdateUserAddressStorage
}

// NewUpdateUserAddressBiz update
func NewUpdateUserAddressBiz(store UpdateUserAddressStorage) *updateUserAddress {
	return &updateUserAddress{store: store}
}

func (biz *updateUserAddress) UpdateUserAddress(id string, v interface{}) error {
	return biz.store.UpdateUserAddress(id, v)
}
