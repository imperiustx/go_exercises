package categorystorage

import (
	"context"

	"github.com/imperiustx/go_excercises/module/category/categorymodel"
)

func (s *sqlStore) FindCategory(
	ctx context.Context, 
	conditions map[string]interface{},
	moreInfo ...string) (*categorymodel.Category, error) {

	db := s.db.Table(categorymodel.Category{}.TableName())

	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}
	
	var category categorymodel.Category

	if err := db.Where(conditions).First(&category).Error; err != nil {
		return nil, err
	}

	return &category, nil
}
