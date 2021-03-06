package categorybusiness

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/category/categorymodel"
)

// CreateCategoryStorage create
type CreateCategoryStorage interface {
	FindCategory(
		ctx context.Context, 
		conditions map[string]interface{},
		moreInfo ...string) (*categorymodel.Category, error)
	CreateCategory(ctx context.Context, data *categorymodel.CategoryCreate) error
}

type createCategory struct {
	store CreateCategoryStorage
}

// NewCreateCategoryBiz create
func NewCreateCategoryBiz(store CreateCategoryStorage) *createCategory {
	return &createCategory{store: store}
}

func (biz *createCategory) CreateNewCategory(ctx context.Context, data *categorymodel.CategoryCreate) error {
	category, err := biz.store.FindCategory(ctx, map[string]interface{}{"name": data.Name})

	if category != nil {
		return common.ErrEntityExisted(categorymodel.EntityName, err)
	}
	
	if err := biz.store.CreateCategory(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(categorymodel.EntityName, err)
	}

	return nil
}
