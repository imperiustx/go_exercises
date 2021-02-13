package restaurantfoodstorage

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/restaurantfood/restaurantfoodmodel"
)

func (s *sqlStore) ListRestaurantFood(
	ctx context.Context,
	paging *common.Paging,
	order *common.OrderSort,
	moreKeys ...string) ([]restaurantfoodmodel.RestaurantFood, error) {

	db := s.db.Table(restaurantfoodmodel.RestaurantFood{}.TableName())
	var restaurantfoods []restaurantfoodmodel.RestaurantFood

	db = db.Where("status not in (0)")

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	db = db.Limit(paging.Limit)
	//FIXME: fix id
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

	if err := db.Find(&restaurantfoods).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return restaurantfoods, nil
}
