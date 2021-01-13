package paymentmethodbusiness

// DeletePaymentMethodStorage delete
type DeletePaymentMethodStorage interface {
	DeletePaymentMethod(id string) error
}

type deletePaymentMethod struct {
	store DeletePaymentMethodStorage
}

// NewDeletePaymentMethodBiz delete
func NewDeletePaymentMethodBiz(store DeletePaymentMethodStorage) *deletePaymentMethod {
	return &deletePaymentMethod{store: store}
}

func (biz *deletePaymentMethod) DeletePaymentMethod(id string) error {
	return biz.store.DeletePaymentMethod(id)
}
