package userbusiness

import "github.com/imperiustx/go_excercises/module/user/usermodel"

// ListUserStorage list
type ListUserStorage interface {
	ListUser() ([]usermodel.User, error)
}

type listUser struct {
	store ListUserStorage
}

// NewListUserBiz list
func NewListUserBiz(store ListUserStorage) *listUser {
	return &listUser{store: store}
}

func (biz *listUser) ListAllUser() ([]usermodel.User, error) {
	return biz.store.ListUser()
}
