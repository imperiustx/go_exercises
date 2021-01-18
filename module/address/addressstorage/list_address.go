package addressstorage

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/address/addressmodel"
)

func (s *sqlStore) ListAddress(ctx context.Context, paging *common.Paging) ([]addressmodel.Address, error) {
	db := s.db
	var addresses []addressmodel.Address

	db = db.Table(addressmodel.Address{}.TableName()).Where("status not in (0)")

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
	if err := db.Order("id asc").Find(&addresses).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return addresses, nil
}
