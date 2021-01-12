package userbusiness

import "github.com/imperiustx/go_excercises/module/user/usermodel"

// GetUserStorage get
type GetUserStorage interface {
	GetUser(id string) (usermodel.User, error)
}

type getUser struct {
	store GetUserStorage
}

// NewGetUserBiz get
func NewGetUserBiz(store GetUserStorage) *getUser {
	return &getUser{store: store}
}

func (biz *getUser) GetUser(id string) (usermodel.User, error) {
	return biz.store.GetUser(id)
}
