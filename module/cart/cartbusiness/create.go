package cartbusiness

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/cart/cartmodel"
)

type CreateStorage interface {
	CreateCart(ctx context.Context, data *cartmodel.CartCreate) error
}

type createBusiness struct {
	createStorage CreateStorage
}

func NewCreateBusiness(createStorage CreateStorage) *createBusiness {
	return &createBusiness{
		createStorage: createStorage,
	}
}

func (business *createBusiness) Create(ctx context.Context, data *cartmodel.CartCreate) error {

	if err := business.createStorage.CreateCart(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(cartmodel.EntityName, err)
	}

	return nil
}
