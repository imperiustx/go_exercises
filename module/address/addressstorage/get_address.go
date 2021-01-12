package addressstorage

import "github.com/imperiustx/go_excercises/module/address/addressmodel"

func (s *sqlStore) GetAddress(id string) (addressmodel.Address, error) {
	db := s.db
	var address addressmodel.Address

	/*

		TODO: convert below to gorm query
			SELECT full_name, email, phone_number, full_address, bank_name, number
			FROM addresss u
			INNER JOIN address_addresses ON u.id = address_addresses.id
			INNER JOIN addresses a ON address_addresses.address_id = a.id
			INNER JOIN address_payment_methods ON u.id = address_payment_methods.address_id
			INNER JOIN payment_methods p ON address_payment_methods.id = p.id
			WHERE u.id = 1;
	*/

	if err := db.Where("id = ?", id).Select("full_address").Find(&address).Error; err != nil {
		return addressmodel.Address{}, err
	}

	return address, nil
}

