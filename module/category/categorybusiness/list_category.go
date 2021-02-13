package categorybusiness

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/category/categorymodel"
)

// ListCategoryStorage list
type ListCategoryStorage interface {
	ListCategory(
		ctx context.Context,
		paging *common.Paging,
		order *common.OrderSort,
		moreKeys ...string) ([]categorymodel.Category, error)
}

type listCategory struct {
	store ListCategoryStorage
}

// NewListCategoryBiz list
func NewListCategoryBiz(store ListCategoryStorage) *listCategory {
	return &listCategory{store: store}
}

func (biz *listCategory) ListAllCategory(
	ctx context.Context,
	paging *common.Paging,
	order *common.OrderSort,
	moreKeys ...string) ([]categorymodel.Category, error) {

	data, err := biz.store.ListCategory(ctx, paging, order)
	if err != nil {
		return nil, common.ErrCannotListEntity(categorymodel.EntityName, err)
	}

	return data, nil
}
