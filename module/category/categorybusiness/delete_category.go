package categorybusiness

import (
	"context"
	"errors"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/category/categorymodel"
)

// DeleteCategoryStorage delete
type DeleteCategoryStorage interface {
	FindCategory(
		ctx context.Context, 
		conditions map[string]interface{},
		moreInfo ...string) (*categorymodel.Category, error)
	DeleteCategory(ctx context.Context, conditions map[string]interface{}) error
}

type deleteCategory struct {
	store DeleteCategoryStorage
}

// NewDeleteCategoryBiz delete
func NewDeleteCategoryBiz(store DeleteCategoryStorage) *deleteCategory {
	return &deleteCategory{store: store}
}

func (biz *deleteCategory) DeleteCategory(ctx context.Context, conditions map[string]interface{}) error {
	category, err := biz.store.FindCategory(ctx, conditions)
	if err != nil {
		return common.ErrCannotGetEntity(categorymodel.EntityName, err)
	}

	if category.Status == 0 {
		return common.ErrCannotGetEntity(categorymodel.EntityName, errors.New("category not found"))
	}

	if err := biz.store.DeleteCategory(ctx, conditions); err != nil {
		return err
	}

	return nil
}
