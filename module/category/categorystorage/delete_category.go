package categorystorage

import (
	"context"

	"github.com/imperiustx/go_excercises/module/category/categorymodel"
)

func (s *sqlStore) DeleteCategory(ctx context.Context, conditions map[string]interface{}) error {
	db := s.db.Table(categorymodel.Category{}.TableName())

	if err := db.Where(conditions).Update("status", 0).Error; err != nil {
		return err
	}

	return nil
}
