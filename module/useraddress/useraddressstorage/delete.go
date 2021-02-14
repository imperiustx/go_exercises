package useraddressstorage

import (
	"context"

	"github.com/imperiustx/go_excercises/module/useraddress/useraddressmodel"
)

func (s *sqlStore) DeleteUserAddress(ctx context.Context, conditions map[string]interface{}) error {
	db := s.db.Table(useraddressmodel.UserAddress{}.TableName())

	if err := db.Where(conditions).Update("status", 0).Error; err != nil {
		return err
	}

	return nil
}
