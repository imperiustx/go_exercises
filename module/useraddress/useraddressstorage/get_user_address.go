package useraddressstorage

import "github.com/imperiustx/go_excercises/module/useraddress/useraddressmodel"

func (s *sqlStore) GetUserAddress(id string) (useraddressmodel.UserAddress, error) {
	db := s.db
	var useraddress useraddressmodel.UserAddress

	/*

		TODO: convert below to gorm query
			SELECT full_name, email, phone_number, full_address, bank_name, number
			FROM useraddresss u
			INNER JOIN useraddress_addresses ON u.id = useraddress_addresses.id
			INNER JOIN addresses a ON useraddress_addresses.address_id = a.id
			INNER JOIN useraddress_payment_methods ON u.id = useraddress_payment_methods.useraddress_id
			INNER JOIN payment_methods p ON useraddress_payment_methods.id = p.id
			WHERE u.id = 1;
	*/

	if err := db.Where("id = ?", id).Select("full_name", "email", "phone_number").Find(&useraddress).Error; err != nil {
		return useraddressmodel.UserAddress{}, err
	}

	return useraddress, nil
}
