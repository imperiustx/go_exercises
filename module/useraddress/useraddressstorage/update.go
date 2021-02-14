package useraddressstorage

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/useraddress/useraddressmodel"
)

func (s *sqlStore) UpdateUserAddress(
	ctx context.Context,
	conditions map[string]interface{},
	data *useraddressmodel.UserAddressUpdate) error {
		
	db := s.db.Table(data.TableName())

	if err := db.Where(conditions).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
