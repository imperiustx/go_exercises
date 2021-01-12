package restaurantstorage

import "github.com/imperiustx/go_excercises/module/restaurant/restaurantmodel"

func (s *sqlStore) GetRestaurant(id string) (restaurantmodel.Restaurant, error) {
	db := s.db
	var restaurant restaurantmodel.Restaurant

	/*

		TODO: convert below to gorm query and fix to RESTAURANT
			SELECT full_name, email, phone_number, full_address, bank_name, number
			FROM users u
			INNER JOIN user_addresses ON u.id = user_addresses.id
			INNER JOIN addresses a ON user_addresses.address_id = a.id
			INNER JOIN user_payment_methods ON u.id = user_payment_methods.user_id
			INNER JOIN payment_methods p ON user_payment_methods.id = p.id
			WHERE u.id = 1;
	*/

	if err := db.Where("id = ?", id).Select("name", "phone_number", "price_range").Find(&restaurant).Error; err != nil {
		return restaurantmodel.Restaurant{}, err
	}

	return restaurant, nil
}
