package useraddressstorage

import (
	"context"

	"github.com/imperiustx/go_excercises/module/useraddress/useraddressmodel"
)

func (s *sqlStore) FindUserAddress(
	ctx context.Context,
	conditions map[string]interface{},
	moreInfo ...string) (*useraddressmodel.UserAddress, error) {

	db := s.db.Table(useraddressmodel.UserAddress{}.TableName())

	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}

	var useraddress useraddressmodel.UserAddress

	if err := db.Where(conditions).First(&useraddress).Error; err != nil {
		return nil, err
	}

	return &useraddress, nil
}
