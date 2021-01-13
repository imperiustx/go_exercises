package userpaymentmethodbusiness

import "github.com/imperiustx/go_excercises/module/userpaymentmethod/userpaymentmethodmodel"

// GetUserPaymentMethodStorage get
type GetUserPaymentMethodStorage interface {
	GetUserPaymentMethod(id string) (userpaymentmethodmodel.UserPaymentMethod, error)
}

type getUserPaymentMethod struct {
	store GetUserPaymentMethodStorage
}

// NewGetUserPaymentMethodBiz get
func NewGetUserPaymentMethodBiz(store GetUserPaymentMethodStorage) *getUserPaymentMethod {
	return &getUserPaymentMethod{store: store}
}

func (biz *getUserPaymentMethod) GetUserPaymentMethod(id string) (userpaymentmethodmodel.UserPaymentMethod, error) {
	return biz.store.GetUserPaymentMethod(id)
}
