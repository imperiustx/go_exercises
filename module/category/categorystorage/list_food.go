package categorystorage

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/category/categorymodel"
	"github.com/imperiustx/go_excercises/module/food/foodmodel"
)

func (s *sqlStore) ListFoodByCategoryID(
	ctx context.Context,
	id int,
	paging *common.Paging,
	order *common.OrderSort,
	moreKeys ...string) ([]foodmodel.Food, error) {

	// SELECT foods.id, foods.name, foods.price, categories.name
	// FROM `food_delivery`.`categories`
	// LEFT JOIN foods ON foods.category_id = categories.id
	// WHERE categories.id = 6;

	db := s.db.Table(categorymodel.Category{}.TableName())
	var results []foodmodel.Food

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

	if err := db.
		Select("f.id, f.restaurant_id, f.category_id, f.name, f.description, f.price, f.images, f.status, f.created_at, f.updated_at").
		Joins("left join foods as f on f.category_id = categories.id").
		Where("categories.id = ?", id).Scan(&results).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return results, nil
}
