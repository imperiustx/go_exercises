package userdevicetokenstorage

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/userdevicetoken/userdevicetokenmodel"
)

func (s *sqlStore) FindUserDeviceToken(
	ctx context.Context,
	conditions map[string]interface{},
	moreInfo ...string) (*userdevicetokenmodel.UserDeviceToken, error) {

	db := s.db.Table(userdevicetokenmodel.UserDeviceToken{}.TableName())

	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}

	var userdevicetoken userdevicetokenmodel.UserDeviceToken

	if err := db.Where(conditions).First(&userdevicetoken).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return &userdevicetoken, nil
}
