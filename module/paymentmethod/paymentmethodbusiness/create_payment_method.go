package paymentmethodbusiness

// CreatePaymentMethodStorage create
type CreatePaymentMethodStorage interface {
	CreatePaymentMethod(v interface{}) error
}

type createPaymentMethod struct {
	store CreatePaymentMethodStorage
}

// NewCreatePaymentMethodBiz create
func NewCreatePaymentMethodBiz(store CreatePaymentMethodStorage) *createPaymentMethod {
	return &createPaymentMethod{store: store}
}

func (biz *createPaymentMethod) CreatePaymentMethod(v interface{}) error {
	return biz.store.CreatePaymentMethod(v)
}
