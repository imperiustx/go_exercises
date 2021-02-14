package foodlikestorage

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/foodlike/foodlikemodel"
)

func (s *sqlStore) ListFoodLike(
	ctx context.Context,
	filter *foodlikemodel.Filter,
	paging *common.Paging,
	order *common.OrderSort,
	moreKeys ...string) ([]foodlikemodel.FoodLike, error) {

	db := s.db.Table(foodlikemodel.FoodLike{}.TableName())
	var foodlikes []foodlikemodel.FoodLike

	switch {
	case filter.FoodID != 0:
		db = db.Where("food_id = ?", filter.FoodID)

		if paging.Cursor > 0 {
			db = db.Where("food_id < ?", paging.Cursor)
		} else {
			db = db.Offset((paging.Page - 1) * paging.Limit)
		}

		if o := order; o != nil {
			if o.Order == "asc" {
				db = db.Order("food_id asc")
			}
			if o.Order == "desc" {
				db = db.Order("food_id desc")
			}
		}
	case filter.UserID != 0:
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

	if err := db.Find(&foodlikes).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return foodlikes, nil
}
