package addressbusiness

import (
	"context"
	"errors"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/address/addressmodel"
)

// UpdateAddressStorage update
type UpdateAddressStorage interface {
	FindAddress(ctx context.Context, id int) (*addressmodel.Address, error)
	UpdateAddress(ctx context.Context, id int, data *addressmodel.AddressUpdate) error
}

type updateAddress struct {
	store UpdateAddressStorage
}

// NewUpdateAddressBiz update
func NewUpdateAddressBiz(store UpdateAddressStorage) *updateAddress {
	return &updateAddress{store: store}
}

func (biz *updateAddress) UpdateAddress(ctx context.Context, id int, data *addressmodel.AddressUpdate) error {
	address, err := biz.store.FindAddress(ctx, id)
	if err != nil {
		return common.ErrCannotGetEntity(addressmodel.EntityName, err)
	}

	if address.Status == 0 {
		return common.ErrCannotGetEntity(addressmodel.EntityName, errors.New("address not found"))
	}

	if err := biz.store.UpdateAddress(ctx, address.ID, data); err != nil {
		return common.ErrCannotUpdateEntity(addressmodel.EntityName, err)
	}

	return nil
}
