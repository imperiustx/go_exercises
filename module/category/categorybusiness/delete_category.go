package categorybusiness

import (
	"context"
	"errors"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/category/categorymodel"
)

// DeleteCategoryStorage delete
type DeleteCategoryStorage interface {
	FindCategory(ctx context.Context, id int) (*categorymodel.Category, error)
	DeleteCategory(id int) error
}

type deleteCategory struct {
	store DeleteCategoryStorage
}

// NewDeleteCategoryBiz delete
func NewDeleteCategoryBiz(store DeleteCategoryStorage) *deleteCategory {
	return &deleteCategory{store: store}
}

func (biz *deleteCategory) DeleteCategory(ctx context.Context, id int) error {
	category, err := biz.store.FindCategory(ctx, id)
	if err != nil {
		return common.ErrCannotGetEntity(categorymodel.EntityName, err)
	}

	if category.Status == 0 {
		return common.ErrCannotGetEntity(categorymodel.EntityName, errors.New("category not found"))
	}

	if err := biz.store.DeleteCategory(id); err != nil {
		return err
	}

	return nil
}
