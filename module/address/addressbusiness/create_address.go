package addressbusiness

// CreateAddressStorage create
type CreateAddressStorage interface {
	CreateAddress(v interface{}) error
}

type createAddress struct {
	store CreateAddressStorage
}

// NewCreateAddressBiz create
func NewCreateAddressBiz(store CreateAddressStorage) *createAddress {
	return &createAddress{store: store}
}

func (biz *createAddress) CreateAddress(v interface{}) error {
	return biz.store.CreateAddress(v)
}
