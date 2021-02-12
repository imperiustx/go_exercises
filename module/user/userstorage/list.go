package userstorage

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/user/usermodel"
)

func (s *sqlStore) ListUser(
	ctx context.Context,
	filter *usermodel.Filter,
	paging *common.Paging,
	order *common.OrderSort,
	moreKeys ...string) ([]usermodel.User, error) {
		
	db := s.db.Table(usermodel.User{}.TableName())
	var users []usermodel.User

	db = db.Where("status not in (0)")

	if f := filter; f != nil {
		if f.Email != "" {
			db = db.Where("email = ?", f.Email)
		}
		if f.Phone != "" {
			db = db.Where("phone = ?", f.Phone)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	db = db.Limit(paging.Limit)

	for _, k := range moreKeys {
		db = db.Preload(k)
	}

	if paging.Cursor > 0 {
		db = db.Where("id < ?", paging.Cursor)
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	if o := order; o != nil {
		if o.Order == "asc" {
			db = db.Order("id asc")
		}
		if o.Order == "desc" {
			db = db.Order("id desc")
		}
	}

	if err := db.Find(&users).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return users, nil
}
