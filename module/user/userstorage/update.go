package userstorage

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/user/usermodel"
)

func (s *sqlStore) UpdateUser(ctx context.Context, id int, data *usermodel.UserUpdate) error {
	db := s.db.Table(data.TableName())

	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
