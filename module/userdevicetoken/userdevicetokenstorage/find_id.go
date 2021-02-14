package userdevicetokenstorage

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/userdevicetoken/userdevicetokenmodel"
)

func (s *sqlStore) FindUserDeviceTokenID(
	ctx context.Context,
	uid int,
	os string,
	moreInfo ...string) (*int, error) {

	db := s.db.Table(userdevicetokenmodel.UserDeviceToken{}.TableName())

	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}

	var userdevicetoken userdevicetokenmodel.UserDeviceToken

	if err := db.Where(&userdevicetokenmodel.UserDeviceToken{
		UserID: uid,
		OS:     os,
	}).First(&userdevicetoken).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return &userdevicetoken.ID, nil
}
