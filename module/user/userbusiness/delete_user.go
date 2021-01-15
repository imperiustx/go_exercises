package userbusiness

import (
	"errors"

	"github.com/imperiustx/go_excercises/module/user/usermodel"
)

// DeleteUserStorage delete
type DeleteUserStorage interface {
	FindUser(id int) (*usermodel.User, error)
	DeleteUser(id int) error
}

type deleteUser struct {
	store DeleteUserStorage
}

// NewDeleteUserBiz delete
func NewDeleteUserBiz(store DeleteUserStorage) *deleteUser {
	return &deleteUser{store: store}
}

func (biz *deleteUser) DeleteUser(id int) error {
	user, err := biz.store.FindUser(id)
	if err != nil {
		return err
	}

	if user.Status == 0 {
		return errors.New("user note found")
	}

	if err := biz.store.DeleteUser(id); err != nil {
		return err
	}

	return nil
}
