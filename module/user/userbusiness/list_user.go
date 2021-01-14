package userbusiness

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/user/usermodel"
)

// ListUserStorage list
type ListUserStorage interface {
	ListUser(ctx context.Context, paging *common.Paging) ([]usermodel.User, error)
}

type listUser struct {
	store ListUserStorage
}

// NewListUserBiz list
func NewListUserBiz(store ListUserStorage) *listUser {
	return &listUser{store: store}
}

func (biz *listUser) ListAllUser(ctx context.Context, paging *common.Paging) ([]usermodel.User, error) {
	data, err := biz.store.ListUser(ctx, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(usermodel.EntityName, err)
	}

	return data, nil
}
