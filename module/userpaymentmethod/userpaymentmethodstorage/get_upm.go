package userpaymentmethodstorage

import "github.com/imperiustx/go_excercises/module/userpaymentmethod/userpaymentmethodmodel"

func (s *sqlStore) GetUserPaymentMethod(id string) (userpaymentmethodmodel.UserPaymentMethod, error) {
	db := s.db
	var userpaymentmethod userpaymentmethodmodel.UserPaymentMethod

	/*

		TODO: convert below to gorm query
			SELECT full_name, email, phone_number, full_address, bank_name, number
			FROM userpaymentmethods u
			INNER JOIN userpaymentmethod_addresses ON u.id = userpaymentmethod_addresses.id
			INNER JOIN addresses a ON userpaymentmethod_addresses.address_id = a.id
			INNER JOIN userpaymentmethod_payment_methods ON u.id = userpaymentmethod_payment_methods.userpaymentmethod_id
			INNER JOIN payment_methods p ON userpaymentmethod_payment_methods.id = p.id
			WHERE u.id = 1;
	*/

	if err := db.Where("id = ?", id).Select("full_name", "email", "phone_number").Find(&userpaymentmethod).Error; err != nil {
		return userpaymentmethodmodel.UserPaymentMethod{}, err
	}

	return userpaymentmethod, nil
}
