package userbusiness

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/user/usermodel"
)

// GetUserStorage get
type GetUserStorage interface {
	FindUser(ctx context.Context, id int) (*usermodel.User, error)
}

type getUserBiz struct {
	store GetUserStorage
}

// NewGetUserBiz get
func NewGetUserBiz(store GetUserStorage) *getUserBiz {
	return &getUserBiz{store: store}
}

func (biz *getUserBiz) GetUser(ctx context.Context, id int) (*usermodel.User, error) {
	user, err := biz.store.FindUser(ctx, id)
	if err != nil {
		return nil, common.ErrCannotGetEntity(usermodel.EntityName, err)
	}

	return user, nil
}
