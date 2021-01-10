package userbusiness

// CreateUserStorage create
type CreateUserStorage interface {
	CreateUser(v interface{}) error
}

type createUser struct {
	store CreateUserStorage
}

// NewCreateUserBiz create
func NewCreateUserBiz(store CreateUserStorage) *createUser {
	return &createUser{store: store}
}

func (biz *createUser) CreateUser(v interface{}) error {
	return biz.store.CreateUser(v)
}
