package foodstorage

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/food/foodmodel"
	"github.com/imperiustx/go_excercises/module/restaurant/restaurantmodel"
)

func (s *sqlStore) ListFoodByRestaurantID(
	ctx context.Context,
	id int,
	paging *common.Paging,
	order *common.OrderSort,
	moreKeys ...string) ([]foodmodel.Food, error) {

	db := s.db.Table(restaurantmodel.Restaurant{}.TableName())

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

	rows, err := db.Where("restaurants.id = ?", id).
		Select("f.id, f.restaurant_id, f.category_id, f.name, f.description, f.price, f.images, f.status, f.created_at, f.updated_at").
		Joins("LEFT JOIN foods AS f ON f.restaurant_id = restaurants.id").Rows()
	if err != nil {
		return nil, common.ErrDB(err)
	}

	var results []foodmodel.Food

	for rows.Next() {
		var result foodmodel.Food

		if err := s.db.ScanRows(rows, &result); err != nil {
			return nil, common.ErrDB(err)
		}
		results = append(results, result)
	}

	return results, nil
}
