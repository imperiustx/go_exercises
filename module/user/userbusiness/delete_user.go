package userbusiness

import (
	"context"
	"errors"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/user/usermodel"
)

// DeleteUserStorage delete
type DeleteUserStorage interface {
	FindUser(ctx context.Context, id int) (*usermodel.User, error)
	DeleteUser(id int) error
}

type deleteUser struct {
	store DeleteUserStorage
}

// NewDeleteUserBiz delete
func NewDeleteUserBiz(store DeleteUserStorage) *deleteUser {
	return &deleteUser{store: store}
}

func (biz *deleteUser) DeleteUser(ctx context.Context, id int) error {
	user, err := biz.store.FindUser(ctx, id)
	if err != nil {
		return common.ErrCannotGetEntity(usermodel.EntityName, err)
	}

	if user.Status == 0 {
		return common.ErrCannotGetEntity(usermodel.EntityName, errors.New("user not found"))
	}

	if err := biz.store.DeleteUser(id); err != nil {
		return err
	}

	return nil
}
