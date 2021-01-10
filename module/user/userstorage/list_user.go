package userstorage

import "github.com/imperiustx/go_excercises/module/user/usermodel"

func (s *sqlStore) ListUser() ([]usermodel.User, error) {
	db := s.db
	var users []usermodel.User

	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
