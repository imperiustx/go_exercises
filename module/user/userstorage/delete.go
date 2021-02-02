package userstorage

import "github.com/imperiustx/go_excercises/module/user/usermodel"

func (s *sqlStore) DeleteUser(conditions map[string]interface{}) error {
	db := s.db.Table(usermodel.User{}.TableName())

	if err := db.Where("id = ?", conditions["id"]).
		Update("status", 0).Error; err != nil {
		return err
	}

	return nil
}
