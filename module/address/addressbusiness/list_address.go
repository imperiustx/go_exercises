package addressbusiness

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/address/addressmodel"
)

// ListAddressStorage list
type ListAddressStorage interface {
	ListAddress(ctx context.Context, paging *common.Paging) ([]addressmodel.Address, error)
}

type listAddress struct {
	store ListAddressStorage
}

// NewListAddressBiz list
func NewListAddressBiz(store ListAddressStorage) *listAddress {
	return &listAddress{store: store}
}

func (biz *listAddress) ListAllAddress(ctx context.Context, paging *common.Paging) ([]addressmodel.Address, error) {
	data, err := biz.store.ListAddress(ctx, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(addressmodel.EntityName, err)
	}

	return data, nil
}
