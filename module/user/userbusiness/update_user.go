package userbusiness

// UpdateUserStorage update
type UpdateUserStorage interface {
	UpdateUser(id string, v interface{}) error
}

type updateUser struct {
	store UpdateUserStorage
}

// NewUpdateUserBiz update
func NewUpdateUserBiz(store UpdateUserStorage) *updateUser {
	return &updateUser{store: store}
}

func (biz *updateUser) UpdateUser(id string, v interface{}) error {
	return biz.store.UpdateUser(id, v)
}
