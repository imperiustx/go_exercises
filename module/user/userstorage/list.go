package userstorage

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/user/usermodel"
)

func (s *sqlStore) ListUser(ctx context.Context, paging *common.Paging) ([]usermodel.User, error) {
	db := s.db
	var users []usermodel.User

	db = db.Table(usermodel.User{}.TableName()).Where("status not in (0)")

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	db = db.Limit(paging.Limit)

	if paging.Cursor > 0 {
		db = db.Where("id < ?", paging.Cursor)
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	// id desc
	if err := db.Order("id asc").Find(&users).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return users, nil
}
