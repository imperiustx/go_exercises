package categorystorage

import (
	"context"

	"github.com/imperiustx/go_excercises/module/category/categorymodel"
)

func (s *sqlStore) FindCategory(ctx context.Context, id int) (*categorymodel.Category, error) {
	db := s.db
	var category categorymodel.Category

	if err := db.Where("id = ?", id).
		First(&category).Error; err != nil {
		return nil, err
	}

	return &category, nil
}
