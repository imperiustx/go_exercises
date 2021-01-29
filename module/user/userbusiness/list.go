package userbusiness

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/user/usermodel"
)

// ListUserStorage list
type ListUserStorage interface {
	ListUser(
		ctx context.Context,
		filter *usermodel.Filter,
		paging *common.Paging,
		order *common.OrderSort,
		moreKeys ...string) ([]usermodel.User, error)
}

type listUser struct {
	store ListUserStorage
}

// NewListUserBiz list
func NewListUserBiz(store ListUserStorage) *listUser {
	return &listUser{store: store}
}

func (biz *listUser) ListAllUser(
	ctx context.Context,
	filter *usermodel.Filter,
	paging *common.Paging,
	order *common.OrderSort) ([]usermodel.User, error) {
	data, err := biz.store.ListUser(ctx, filter, paging, order)
	if err != nil {
		return nil, common.ErrCannotListEntity(usermodel.EntityName, err)
	}

	return data, nil
}
