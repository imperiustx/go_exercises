package userbusiness

import (
	"context"
	"errors"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/user/usermodel"
)

// DeleteUserStorage delete
type DeleteUserStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
	DeleteUser(conditions map[string]interface{}) error
}

type deleteUser struct {
	store     DeleteUserStorage
	requester common.Requester
}

// NewDeleteUserBiz delete
func NewDeleteUserBiz(store DeleteUserStorage, requester common.Requester) *deleteUser {
	return &deleteUser{
		store:     store,
		requester: requester,
	}
}

func (biz *deleteUser) DeleteUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) error {
	user, err := biz.store.FindUser(ctx, conditions)
	if err != nil {
		return common.ErrCannotGetEntity(usermodel.EntityName, err)
	}

	if user.Status == 0 {
		return common.ErrCannotGetEntity(usermodel.EntityName, errors.New("user not found"))
	}

	isAuthor := biz.requester.GetUserId() == user.ID
	isAdmin := biz.requester.GetRole() == "admin"

	if !isAuthor && !isAdmin {
		return common.ErrNoPermission(nil)
	}

	if err := biz.store.DeleteUser(conditions); err != nil {
		return common.ErrCannotDeleteEntity(usermodel.EntityName, err)
	}

	return nil
}
