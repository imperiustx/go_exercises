package restaurantlikestorage

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/restaurantlike/restaurantlikemodel"
)

func (s *sqlStore) ListRestaurantLike(ctx context.Context, paging *common.Paging) ([]restaurantlikemodel.RestaurantLike, error) {
	db := s.db
	var restaurantlikes []restaurantlikemodel.RestaurantLike

	db = db.Table(restaurantlikemodel.RestaurantLike{}.TableName()).Where("status not in (0)")

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
	if err := db.Order("id asc").Find(&restaurantlikes).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return restaurantlikes, nil
}
