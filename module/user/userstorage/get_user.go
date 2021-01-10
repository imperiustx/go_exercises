package userstorage

import "github.com/imperiustx/go_excercises/module/user/usermodel"

func (s *sqlStore) GetUser(id string) (usermodel.User, error) {
	db := s.db
	var user usermodel.User

	/*

		TODO: convert below to gorm query
			SELECT full_name, email, phone_number, full_address, bank_name, number
			FROM users u
			INNER JOIN user_addresses ON u.id = user_addresses.id
			INNER JOIN addresses a ON user_addresses.address_id = a.id
			INNER JOIN user_payment_methods ON u.id = user_payment_methods.user_id
			INNER JOIN payment_methods p ON user_payment_methods.id = p.id
			WHERE u.id = 1;
	*/

	if err := db.Where("id = ?", id).Select("full_name", "email", "phone_number").Find(&user).Error; err != nil {
		return usermodel.User{}, err
	}

	return user, nil
}
