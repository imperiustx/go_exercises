package userdevicetokenstorage

import (
	"context"

	"github.com/imperiustx/go_excercises/module/userdevicetoken/userdevicetokenmodel"
)

func (s *sqlStore) DeleteUserDeviceToken(
	ctx context.Context,
	conditions map[string]interface{}) error {
		
	db := s.db.Table(userdevicetokenmodel.UserDeviceToken{}.TableName())

	if err := db.Where(conditions).Update("status", 0).Error; err != nil {
		return err
	}

	return nil
}
