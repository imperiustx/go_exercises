package citybusiness

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/city/citymodel"
)

// ListCityStorage list
type ListCityStorage interface {
	ListCity(ctx context.Context, paging *common.Paging) ([]citymodel.City, error)
}

type listCity struct {
	store ListCityStorage
}

// NewListCityBiz list
func NewListCityBiz(store ListCityStorage) *listCity {
	return &listCity{store: store}
}

func (biz *listCity) ListAllCity(ctx context.Context, paging *common.Paging) ([]citymodel.City, error) {
	data, err := biz.store.ListCity(ctx, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(citymodel.EntityName, err)
	}

	return data, nil
}
