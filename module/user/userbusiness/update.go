package userbusiness

import (
	"context"
	"errors"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/user/usermodel"
)

// UpdateUserStorage update
type UpdateUserStorage interface {
	FindUser(
		ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string) (*usermodel.User, error)
	UpdateUser(
		ctx context.Context,
		conditions map[string]interface{},
		data *usermodel.UserUpdate) error
}

type updateUser struct {
	store     UpdateUserStorage
	requester common.Requester
}

// NewUpdateUserBiz update
func NewUpdateUserBiz(store UpdateUserStorage, requester common.Requester) *updateUser {
	return &updateUser{
		store:     store,
		requester: requester,
	}
}

func (biz *updateUser) UpdateUser(
	ctx context.Context,
	conditions map[string]interface{},
	data *usermodel.UserUpdate) error {
	user, err := biz.store.FindUser(ctx, conditions)
	if err != nil {
		return common.ErrCannotGetEntity(usermodel.EntityName, err)
	}

	if user.Status == 0 {
		return common.ErrCannotGetEntity(usermodel.EntityName, errors.New("user not found"))
	}

	if biz.requester.GetUserId() != user.ID {
		return common.ErrNoPermission(nil)
	}

	if err := biz.store.UpdateUser(ctx, conditions, data); err != nil {
		return common.ErrCannotUpdateEntity(usermodel.EntityName, err)
	}

	return nil
}
