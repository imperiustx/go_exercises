package userpaymentmethodbusiness

// CreateUserPaymentMethodStorage create
type CreateUserPaymentMethodStorage interface {
	CreateUserPaymentMethod(v interface{}) error
}

type createUserPaymentMethod struct {
	store CreateUserPaymentMethodStorage
}

// NewCreateUserPaymentMethodBiz create
func NewCreateUserPaymentMethodBiz(store CreateUserPaymentMethodStorage) *createUserPaymentMethod {
	return &createUserPaymentMethod{store: store}
}

func (biz *createUserPaymentMethod) CreateUserPaymentMethod(v interface{}) error {
	return biz.store.CreateUserPaymentMethod(v)
}
