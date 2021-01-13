package paymentmethodstorage

import "github.com/imperiustx/go_excercises/module/paymentmethod/paymentmethodmodel"

func (s *sqlStore) ListPaymentMethod() ([]paymentmethodmodel.PaymentMethod, error) {
	db := s.db
	var paymentMethods []paymentmethodmodel.PaymentMethod

	if err := db.Find(&paymentMethods).Error; err != nil {
		return nil, err
	}

	return paymentMethods, nil
}
