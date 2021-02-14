package userdevicetokenstorage

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/userdevicetoken/userdevicetokenmodel"
)

func (s *sqlStore) UpdateUserDeviceToken(
	ctx context.Context,
	conditions map[string]interface{},
	data *userdevicetokenmodel.UserDeviceTokenUpdate) error {
		
	db := s.db.Table(data.TableName())

	if err := db.Where(conditions).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
