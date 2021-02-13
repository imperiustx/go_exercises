package userstorage

import (
	"context"

	"github.com/imperiustx/go_excercises/module/user/usermodel"
)

func (s *sqlStore) DeleteUser(
	ctx context.Context, 
	conditions map[string]interface{}) error {
	db := s.db.Table(usermodel.User{}.TableName())

	if err := db.Where(conditions).Update("status", 0).Error; err != nil {
		return err
	}

	return nil
}
