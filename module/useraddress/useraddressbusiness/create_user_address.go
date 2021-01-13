package useraddressbusiness

// CreateUserAddressStorage create
type CreateUserAddressStorage interface {
	CreateUserAddress(v interface{}) error
}

type createUserAddress struct {
	store CreateUserAddressStorage
}

// NewCreateUserAddressBiz create
func NewCreateUserAddressBiz(store CreateUserAddressStorage) *createUserAddress {
	return &createUserAddress{store: store}
}

func (biz *createUserAddress) CreateUserAddress(v interface{}) error {
	return biz.store.CreateUserAddress(v)
}
