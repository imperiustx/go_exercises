package userbusiness

import (
	"context"
	"errors"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/user/usermodel"
)

// UpdateUserStorage update
type UpdateUserStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
	UpdateUser(ctx context.Context, id int, data *usermodel.UserUpdate) error
}

type updateUser struct {
	store UpdateUserStorage
}

// NewUpdateUserBiz update
func NewUpdateUserBiz(store UpdateUserStorage) *updateUser {
	return &updateUser{store: store}
}

func (biz *updateUser) UpdateUser(ctx context.Context, id int, data *usermodel.UserUpdate) error {
	user, err := biz.store.FindUser(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return common.ErrCannotGetEntity(usermodel.EntityName, err)
	}

	if user.Status == 0 {
		return common.ErrCannotGetEntity(usermodel.EntityName, errors.New("user not found"))
	}

	if err := biz.store.UpdateUser(ctx, user.ID, data); err != nil {
		return common.ErrCannotUpdateEntity(usermodel.EntityName, err)
	}

	return nil
}
