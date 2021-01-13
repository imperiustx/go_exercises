package userpaymentmethodbusiness

import "github.com/imperiustx/go_excercises/module/userpaymentmethod/userpaymentmethodmodel"

// ListUserPaymentMethodStorage list
type ListUserPaymentMethodStorage interface {
	ListUserPaymentMethod() ([]userpaymentmethodmodel.UserPaymentMethod, error)
}

type listUserPaymentMethod struct {
	store ListUserPaymentMethodStorage
}

// NewListUserPaymentMethodBiz list
func NewListUserPaymentMethodBiz(store ListUserPaymentMethodStorage) *listUserPaymentMethod {
	return &listUserPaymentMethod{store: store}
}

func (biz *listUserPaymentMethod) ListAllUserPaymentMethod() ([]userpaymentmethodmodel.UserPaymentMethod, error) {
	return biz.store.ListUserPaymentMethod()
}
