package addressbusiness

import (
	"context"
	"errors"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/address/addressmodel"
)

// DeleteAddressStorage delete
type DeleteAddressStorage interface {
	FindAddress(ctx context.Context, id int) (*addressmodel.Address, error)
	DeleteAddress(id int) error
}

type deleteAddress struct {
	store DeleteAddressStorage
}

// NewDeleteAddressBiz delete
func NewDeleteAddressBiz(store DeleteAddressStorage) *deleteAddress {
	return &deleteAddress{store: store}
}

func (biz *deleteAddress) DeleteAddress(ctx context.Context, id int) error {
	address, err := biz.store.FindAddress(ctx, id)
	if err != nil {
		return common.ErrCannotGetEntity(addressmodel.EntityName, err)
	}

	if address.Status == 0 {
		return common.ErrCannotGetEntity(addressmodel.EntityName, errors.New("address not found"))
	}

	if err := biz.store.DeleteAddress(id); err != nil {
		return err
	}

	return nil
}
