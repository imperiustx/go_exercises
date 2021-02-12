package useraddressstorage

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/useraddress/useraddressmodel"
)

func (s *sqlStore) ListUserAddress(
	ctx context.Context,
	filter *useraddressmodel.Filter,
	paging *common.Paging,
	order *common.OrderSort,
	moreKeys ...string) ([]useraddressmodel.UserAddress, error) {

	db := s.db.Table(useraddressmodel.UserAddress{}.TableName())
	var useraddresses []useraddressmodel.UserAddress

	db = db.Where("status not in (0)")

	if f := filter; f != nil {
		if f.CityID != 0 {
			db = db.Where("city_id = ?", f.CityID)
		}
		if f.UserID != 0 {
			db = db.Where("user_id = ?", f.UserID)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	db = db.Limit(paging.Limit)

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

	if err := db.Find(&useraddresses).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return useraddresses, nil
}
