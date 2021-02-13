package categorybusiness

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/category/categorymodel"
)

// GetCategoryStorage get
type GetCategoryStorage interface {
	FindCategory(
		ctx context.Context, 
		conditions map[string]interface{},
		moreInfo ...string) (*categorymodel.Category, error)
}

type getCategoryBiz struct {
	store GetCategoryStorage
}

// NewGetCategoryBiz get
func NewGetCategoryBiz(store GetCategoryStorage) *getCategoryBiz {
	return &getCategoryBiz{store: store}
}

func (biz *getCategoryBiz) GetCategory(
	ctx context.Context, 
	conditions map[string]interface{},
	moreInfo ...string) (*categorymodel.Category, error) {

	category, err := biz.store.FindCategory(ctx, conditions)
	if err != nil {
		return nil, common.ErrCannotGetEntity(categorymodel.EntityName, err)
	}

	return category, nil
}
