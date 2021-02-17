package restaurantlikestorage

import (
	"context"
	"errors"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/restaurantlike/restaurantlikemodel"
)

func (s *sqlStore) ListRestaurantLike(
	ctx context.Context,
	filter *restaurantlikemodel.Filter,
	paging *common.Paging,
	order *common.OrderSort,
	moreKeys ...string) ([]restaurantlikemodel.RestaurantLike, error) {

	db := s.db.Table(restaurantlikemodel.RestaurantLike{}.TableName())
	var restaurantlikes []restaurantlikemodel.RestaurantLike

	switch {
	case filter.RestaurantID != "" && filter.UserID != "":
		return nil, common.ErrCannotList(errors.New("Both restaurant_id and user_id is chosen"))
		
	case filter.RestaurantID != "":
		db = db.Where("restaurant_id = ?", filter.RestaurantID)

		if paging.Cursor > 0 {
			db = db.Where("restaurant_id < ?", paging.Cursor)
		} else {
			db = db.Offset((paging.Page - 1) * paging.Limit)
		}

		if o := order; o != nil {
			if o.Order == "asc" {
				db = db.Order("restaurant_id asc")
			}
			if o.Order == "desc" {
				db = db.Order("restaurant_id desc")
			}
		}

	case filter.UserID != "":
		db = db.Where("user_id = ?", filter.UserID)
		if paging.Cursor > 0 {
			db = db.Where("user_id < ?", paging.Cursor)
		} else {
			db = db.Offset((paging.Page - 1) * paging.Limit)
		}

		if o := order; o != nil {
			if o.Order == "asc" {
				db = db.Order("user_id asc")
			}
			if o.Order == "desc" {
				db = db.Order("user_id desc")
			}
		}

	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	db = db.Limit(paging.Limit)

	if err := db.Find(&restaurantlikes).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return restaurantlikes, nil
}
