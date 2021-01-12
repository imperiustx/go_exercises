package userbusiness

// DeleteUserStorage delete
type DeleteUserStorage interface {
	DeleteUser(id string) error
}

type deleteUser struct {
	store DeleteUserStorage
}

// NewDeleteUserBiz delete
func NewDeleteUserBiz(store DeleteUserStorage) *deleteUser {
	return &deleteUser{store: store}
}

func (biz *deleteUser) DeleteUser(id string) error {
	return biz.store.DeleteUser(id)
}
