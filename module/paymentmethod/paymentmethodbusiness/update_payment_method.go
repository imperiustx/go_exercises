package paymentmethodbusiness

// UpdatePaymentMethodStorage update
type UpdatePaymentMethodStorage interface {
	UpdatePaymentMethod(id string, v interface{}) error
}

type updatePaymentMethod struct {
	store UpdatePaymentMethodStorage
}

// NewUpdatePaymentMethodBiz update
func NewUpdatePaymentMethodBiz(store UpdatePaymentMethodStorage) *updatePaymentMethod {
	return &updatePaymentMethod{store: store}
}

func (biz *updatePaymentMethod) UpdatePaymentMethod(id string, v interface{}) error {
	return biz.store.UpdatePaymentMethod(id, v)
}
