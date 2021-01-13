package paymentmethodstorage

import "github.com/imperiustx/go_excercises/module/paymentmethod/paymentmethodmodel"

func (s *sqlStore) GetPaymentMethod(id string) (paymentmethodmodel.PaymentMethod, error) {
	db := s.db
	var paymentMethod paymentmethodmodel.PaymentMethod

	/*

		TODO: convert below to gorm query
			SELECT full_name, email, phone_number, full_address, bank_name, number
			FROM paymentMethods u
			INNER JOIN paymentMethod_addresses ON u.id = paymentMethod_addresses.id
			INNER JOIN addresses a ON paymentMethod_addresses.address_id = a.id
			INNER JOIN paymentMethod_payment_methods ON u.id = paymentMethod_payment_methods.paymentMethod_id
			INNER JOIN payment_methods p ON paymentMethod_payment_methods.id = p.id
			WHERE u.id = 1;
	*/

	if err := db.Where("id = ?", id).Select("bank_name", "paymentmethod_name", "number").Find(&paymentMethod).Error; err != nil {
		return paymentmethodmodel.PaymentMethod{}, err
	}

	return paymentMethod, nil
}
