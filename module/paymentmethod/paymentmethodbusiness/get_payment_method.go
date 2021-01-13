package paymentmethodbusiness

import "github.com/imperiustx/go_excercises/module/paymentmethod/paymentmethodmodel"

// GetPaymentMethodStorage get
type GetPaymentMethodStorage interface {
	GetPaymentMethod(id string) (paymentmethodmodel.PaymentMethod, error)
}

type getPaymentMethod struct {
	store GetPaymentMethodStorage
}

// NewGetPaymentMethodBiz get
func NewGetPaymentMethodBiz(store GetPaymentMethodStorage) *getPaymentMethod {
	return &getPaymentMethod{store: store}
}

func (biz *getPaymentMethod) GetPaymentMethod(id string) (paymentmethodmodel.PaymentMethod, error) {
	return biz.store.GetPaymentMethod(id)
}
