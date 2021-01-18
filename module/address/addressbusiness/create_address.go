package addressbusiness

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/address/addressmodel"
)

// CreateAddressStorage create
type CreateAddressStorage interface {
	CreateAddress(ctx context.Context, data *addressmodel.AddressCreate) error
}

type createAddress struct {
	store CreateAddressStorage
}

// NewCreateAddressBiz create
func NewCreateAddressBiz(store CreateAddressStorage) *createAddress {
	return &createAddress{store: store}
}

func (biz *createAddress) CreateNewAddress(ctx context.Context, data *addressmodel.AddressCreate) error {
	if err := biz.store.CreateAddress(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(addressmodel.EntityName, err)
	}

	return nil
}
