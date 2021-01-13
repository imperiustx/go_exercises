package userpaymentmethodbusiness

// DeleteUserPaymentMethodStorage delete
type DeleteUserPaymentMethodStorage interface {
	DeleteUserPaymentMethod(id string) error
}

type deleteUserPaymentMethod struct {
	store DeleteUserPaymentMethodStorage
}

// NewDeleteUserPaymentMethodBiz delete
func NewDeleteUserPaymentMethodBiz(store DeleteUserPaymentMethodStorage) *deleteUserPaymentMethod {
	return &deleteUserPaymentMethod{store: store}
}

func (biz *deleteUserPaymentMethod) DeleteUserPaymentMethod(id string) error {
	return biz.store.DeleteUserPaymentMethod(id)
}
