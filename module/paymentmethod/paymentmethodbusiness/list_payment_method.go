package paymentmethodbusiness

import "github.com/imperiustx/go_excercises/module/paymentmethod/paymentmethodmodel"

// ListPaymentMethodStorage list
type ListPaymentMethodStorage interface {
	ListPaymentMethod() ([]paymentmethodmodel.PaymentMethod, error)
}

type listPaymentMethod struct {
	store ListPaymentMethodStorage
}

// NewListPaymentMethodBiz list
func NewListPaymentMethodBiz(store ListPaymentMethodStorage) *listPaymentMethod {
	return &listPaymentMethod{store: store}
}

func (biz *listPaymentMethod) ListAllPaymentMethod() ([]paymentmethodmodel.PaymentMethod, error) {
	return biz.store.ListPaymentMethod()
}
