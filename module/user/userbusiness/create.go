package userbusiness

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/user/usermodel"
)

// CreateUserStorage create
type CreateUserStorage interface {
	CreateUser(ctx context.Context, data *usermodel.UserCreate) error
}

type createUser struct {
	store CreateUserStorage
}

// NewCreateUserBiz create
func NewCreateUserBiz(store CreateUserStorage) *createUser {
	return &createUser{store: store}
}

func (biz *createUser) CreateNewUser(ctx context.Context, data *usermodel.UserCreate) error {
	if err := biz.store.CreateUser(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(usermodel.EntityName, err)
	}

	return nil
}
