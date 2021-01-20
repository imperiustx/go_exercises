package categorybusiness

import (
	"context"
	"errors"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/category/categorymodel"
)

// UpdateCategoryStorage update
type UpdateCategoryStorage interface {
	FindCategory(ctx context.Context, id int) (*categorymodel.Category, error)
	UpdateCategory(ctx context.Context, id int, data *categorymodel.CategoryUpdate) error
}

type updateCategory struct {
	store UpdateCategoryStorage
}

// NewUpdateCategoryBiz update
func NewUpdateCategoryBiz(store UpdateCategoryStorage) *updateCategory {
	return &updateCategory{store: store}
}

func (biz *updateCategory) UpdateCategory(ctx context.Context, id int, data *categorymodel.CategoryUpdate) error {
	category, err := biz.store.FindCategory(ctx, id)
	if err != nil {
		return common.ErrCannotGetEntity(categorymodel.EntityName, err)
	}

	if category.Status == 0 {
		return common.ErrCannotGetEntity(categorymodel.EntityName, errors.New("category not found"))
	}

	if err := biz.store.UpdateCategory(ctx, category.ID, data); err != nil {
		return common.ErrCannotUpdateEntity(categorymodel.EntityName, err)
	}

	return nil
}
