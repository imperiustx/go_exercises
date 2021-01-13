package useraddressbusiness

// DeleteUserAddressStorage delete
type DeleteUserAddressStorage interface {
	DeleteUserAddress(id string) error
}

type deleteUserAddress struct {
	store DeleteUserAddressStorage
}

// NewDeleteUserAddressBiz delete
func NewDeleteUserAddressBiz(store DeleteUserAddressStorage) *deleteUserAddress {
	return &deleteUserAddress{store: store}
}

func (biz *deleteUserAddress) DeleteUserAddress(id string) error {
	return biz.store.DeleteUserAddress(id)
}
