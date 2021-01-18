package addressbusiness

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/address/addressmodel"
)

// GetAddressStorage get
type GetAddressStorage interface {
	FindAddress(ctx context.Context, id int) (*addressmodel.Address, error)
}

type getAddressBiz struct {
	store GetAddressStorage
}

// NewGetAddressBiz get
func NewGetAddressBiz(store GetAddressStorage) *getAddressBiz {
	return &getAddressBiz{store: store}
}

func (biz *getAddressBiz) GetAddress(ctx context.Context, id int) (*addressmodel.Address, error) {
	address, err := biz.store.FindAddress(ctx, id)
	if err != nil {
		return nil, common.ErrCannotGetEntity(addressmodel.EntityName, err)
	}

	return address, nil
}
