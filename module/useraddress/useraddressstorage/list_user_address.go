package useraddressstorage

import "github.com/imperiustx/go_excercises/module/useraddress/useraddressmodel"

func (s *sqlStore) ListUserAddress() ([]useraddressmodel.UserAddress, error) {
	db := s.db
	var useraddresss []useraddressmodel.UserAddress

	if err := db.Find(&useraddresss).Error; err != nil {
		return nil, err
	}

	return useraddresss, nil
}
