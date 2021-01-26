package userbusiness

import (
	"context"
	"errors"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/user/usermodel"
)

// ReactiveUserStorage reactive
type ReactiveUserStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
	ReactiveUser(conditions map[string]interface{}) error
}

type reactiveUser struct {
	store ReactiveUserStorage
}

// NewReactiveUserBiz reactive
func NewReactiveUserBiz(store ReactiveUserStorage) *reactiveUser {
	return &reactiveUser{store: store}
}

func (biz *reactiveUser) ReactiveUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) error {
	user, err := biz.store.FindUser(ctx, conditions)
	if err != nil {
		return common.ErrCannotGetEntity(usermodel.EntityName, err)
	}

	if user.Status != 0 {
		return common.ErrCannotGetEntity(usermodel.EntityName, errors.New("user not deactive"))
	}

	if err := biz.store.ReactiveUser(conditions); err != nil {
		return err
	}

	return nil
}
