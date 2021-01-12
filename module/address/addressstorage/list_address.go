package addressstorage

import "github.com/imperiustx/go_excercises/module/address/addressmodel"


func (s *sqlStore) ListAddress() ([]addressmodel.Address, error) {
	db := s.db
	var addresses []addressmodel.Address

	if err := db.Find(&addresses).Error; err != nil {
		return nil, err
	}

	return addresses, nil
}
