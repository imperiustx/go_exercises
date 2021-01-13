package userpaymentmethodbusiness

// UpdateUserPaymentMethodStorage update
type UpdateUserPaymentMethodStorage interface {
	UpdateUserPaymentMethod(id string, v interface{}) error
}

type updateUserPaymentMethod struct {
	store UpdateUserPaymentMethodStorage
}

// NewUpdateUserPaymentMethodBiz update
func NewUpdateUserPaymentMethodBiz(store UpdateUserPaymentMethodStorage) *updateUserPaymentMethod {
	return &updateUserPaymentMethod{store: store}
}

func (biz *updateUserPaymentMethod) UpdateUserPaymentMethod(id string, v interface{}) error {
	return biz.store.UpdateUserPaymentMethod(id, v)
}
